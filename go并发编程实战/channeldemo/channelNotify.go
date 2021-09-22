package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//信号通知，执行耗时任务时，优雅退出

func main() {
	closing := make(chan struct{})
	closed := make(chan struct{})
	// 模拟业务处理
	go func() {
		for {
			select {
			case <-closing:
				fmt.Println("收到关闭信号通知，退出当前业务处理")
				return
			default:
				// 业务计算
				fmt.Println("working...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	// 处理CTRL+C等中断信号
	termChan := make(chan os.Signal)
	//接收终断信号通知
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	//阻塞收到信号后，退出当前任务，并开始清理工作
	<-termChan
	//退出当前任务
	close(closing)
	//开始清理工作
	go doCleanup(closed)
	select {
	case <-closed:
		fmt.Println("清理工作正常退出!")
	case <-time.After(5*time.Second):
		fmt.Println("清理超时，不等了")
	}
	fmt.Println("程序优雅退出!")
}

func doCleanup(closed chan struct{}) {
	time.Sleep(time.Minute)
	close(closed)
}
