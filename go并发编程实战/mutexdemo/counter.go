package main

import (
	"fmt"
	"sync"
)

// 在这个例子中，我们创建了 10 个 goroutine，同时不断地对一个变量（count）进行加 1 操作，每个 goroutine 负责执行 10 万次的加 1 操作，我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万)。

//没有使用互斥锁mutex,结果错误
//检查data race错误： go run -race "/Users/duanfuhao/learngo/go并发编程实战/mutexdemo/counter.go"
// func main() {
// 	var count = 0
// 	// 使用WaitGroup等待10个goroutine完成
// 	var wg sync.WaitGroup
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			defer wg.Done()
// 			// 对变量count执行10次加1
// 			for j := 0; j < 100000; j++ {
// 				count++
// 			}
// 		}()
// 	}
// 	// 等待10个goroutine完成
// 	wg.Wait()
// 	fmt.Println(count)
// }

//参考文档:https://zhuanlan.zhihu.com/p/365552668
//使用互斥锁,结果正确
func main() {
	var count = 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
