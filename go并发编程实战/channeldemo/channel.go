package main

import (
	"fmt"
	"time"
)

//令牌参数
type token struct{}

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
