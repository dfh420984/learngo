package main

import (
	"fmt"
	"time"
)

//令牌参数
type token struct{}

//有四个 goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
func main() {
	chs := make([]chan token, 4)
	for i := 0; i < 4; i++ {
		chs[i] = make(chan token)
	}
	//创建4个worker
	for i := 0; i < 4; i++ {
		go worker(i, chs[i], chs[(i+1)%4])
	}
	//给第一个chs[0]赋值
	chs[0] <- token{}
	//永久阻塞
	select {}
}

func worker(i int, ch chan token, chNext chan token) {
	for {
		token := <-ch // 取得令牌
		fmt.Println(i + 1)
		time.Sleep(time.Second)
		chNext <- token
	}
}
