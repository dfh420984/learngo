package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 使用读锁保护
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	var counter Counter
	//开启10个读协程
	for i := 0; i < 10; i++ {
		go func() { // 10个reader
			fmt.Println(counter.Count())
			time.Sleep(time.Millisecond)
		}()
	}

	for { // 一个writer
		counter.Incr()
		time.Sleep(time.Second)
	}
}
