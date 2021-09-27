package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

//随机生成一个Config节点
func loadNewConfig() Config {
	return Config{
		NodeName: "北京",
		Addr:     "127.0.0.1",
		Count:    rand.Int31(),
	}
}

// 首先，我们启动一个 goroutine，然后让它随机 sleep 一段时间，之后就变更一下配置，并通过我们前面学到的 Cond 并发原语，通知其它的 reader 去加载新的配置。接下来，我们启动一个 goroutine 等待配置变更的信号，一旦有变更，它就会加载最新的配置。
func main() {
	var config atomic.Value
	var cond = sync.NewCond(&sync.Mutex{})
	//保存出使节点信息
	config.Store(loadNewConfig())
	fmt.Printf("init Config: %+v \n", config.Load())

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() //通知等待者配置已变更
		}
	}()

	//获取新的config
	go func() {
		for {
			cond.L.Lock()
			cond.Wait() // 等待变更信号
			c := config.Load()
			fmt.Printf("new config: %+v \n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}
