package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// sleep 1秒，然后计数值加1
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

//获得当前记数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	var counter Counter
	//开启10个worker线程
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	// 检查点，等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}
