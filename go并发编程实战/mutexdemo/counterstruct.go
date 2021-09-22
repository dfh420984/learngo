package main

import (
	"fmt"
	"sync"
)

//Mutex 会嵌入到其它 struct 中使用
type Counter struct {
	count uint64
	mu    sync.Mutex
}

//把获取锁、释放锁、计数加一的逻辑封装成一个方法,对外不需要暴露锁等逻辑
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}
