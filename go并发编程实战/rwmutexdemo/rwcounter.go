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

//在这个例子中，使用 10 个 goroutine 进行读操作，每读取一次，sleep 1 毫秒，同时，还有一个 gorotine 进行写操作，每一秒写一次，这是一个 1 writer-n reader 的读写场景，而且写操作还不是很频繁（一秒一次）
func main() {
	var counter Counter
	//开启10个读协程
	for i := 0; i < 2; i++ {
		go func() { // 2个reader
			for {
				fmt.Println(counter.Count())
				time.Sleep(time.Second)
			}
		}()
	}

	for { // 一个writer
		counter.Incr()
		time.Sleep(time.Second)
	}
}

//坑点demo,重入，因为读写锁内部基于互斥锁实现对 writer 的并发访问，而互斥锁本身是有重入问题的，writer 重入调用 Lock 的时候，就会出现死锁的现象
// func foo(l *sync.RWMutex) {
//     fmt.Println("in foo")
//     l.Lock()
//     bar(l)
//     l.Unlock()
// }

// func bar(l *sync.RWMutex) {
//     l.Lock()
//     fmt.Println("in bar")
//     l.Unlock()
// }

// func main() {
//     l := &sync.RWMutex{}
//     foo(l)
// }

//坑点demo,递归，这两个 goroutine 互相持有锁并等待，谁也不会退让一步，满足了“writer 依赖活跃的 reader -> 活跃的 reader 依赖新来的 reader -> 新来的 reader 依赖 writer”的死锁条件，所以就导致了死锁的产生
// func main() {
//     var mu sync.RWMutex

//     // writer,稍微等待，然后制造一个调用Lock的场景
//     go func() {
//         time.Sleep(200 * time.Millisecond)
//         mu.Lock()
//         fmt.Println("Lock")
//         time.Sleep(100 * time.Millisecond)
//         mu.Unlock()
//         fmt.Println("Unlock")
//     }()

//     go func() {
//         factorial(&mu, 10) // 计算10的阶乘, 10!
//     }()
    
//     select {}
// }

// // 递归调用计算阶乘
// func factorial(m *sync.RWMutex, n int) int {
//     if n < 1 { // 阶乘退出条件 
//         return 0
//     }
//     fmt.Println("RLock")
//     m.RLock()
//     defer func() {
//         fmt.Println("RUnlock")
//         m.RUnlock()
//     }()
//     time.Sleep(100 * time.Millisecond)
//     return factorial(m, n-1) * n // 递归调用
// }
