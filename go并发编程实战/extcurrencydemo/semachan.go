package main

import (
	"fmt"
	"sync"
	"time"
)

//实现信号量，比较典型的就是使用 Channel 来实现。
// Semaphore 数据结构，并且还实现了Locker接口
type semaphore struct {
	sync.Locker
	ch chan struct{}
}

func NewSemaphore(capacity int) *semaphore {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &semaphore{
		ch: make(chan struct{}, capacity),
	}
}

// 请求一个资源
func (s *semaphore) Lock() {
	s.ch <- struct{}{}
}

// 释放一个资源
func (s *semaphore) Unlock() {
	<-s.ch
}

func (s *semaphore) resource(i int) {
	s.Lock()
	fmt.Printf("start %d \n", i)
	time.Sleep(1 * time.Second) //模拟耗时操作
	fmt.Printf("end %d \n", i)
	defer s.Unlock()
}

func main() {
	sema := NewSemaphore(3)
	defer close(sema.ch)
	for i := 0; i < 5; i++ {
		go sema.resource(i)
	}
	time.Sleep(6 * time.Second)
	fmt.Println("main end")
}
