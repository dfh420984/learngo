package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

//Cond 的使用其实没那么简单。它的复杂在于：一，这段代码有时候需要加锁，有时候可以不加；二，Wait 唤醒后需要检查条件；三，条件变量的更改，其实是需要原子操作或者互斥锁保护的。所以，有的开发者会认为，Cond 是唯一难以掌握的 Go 并发原语。
func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int
	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			// 加锁更改等待条件
			c.L.Lock()
			defer c.L.Unlock()
			ready++
			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()
	//所有的运动员是否就绪
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}

// 使用 Cond 的 2 个常见错误
// Cond 最常见的使用错误，也就是调用 Wait 的时候没有加锁
// func main() {
//     c := sync.NewCond(&sync.Mutex{})
//     var ready int

//     for i := 0; i < 10; i++ {
//         go func(i int) {
//             time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

//             // 加锁更改等待条件
//             c.L.Lock()
//             ready++
//             c.L.Unlock()

//             log.Printf("运动员#%d 已准备就绪\n", i)
//             // 广播唤醒所有的等待者
//             c.Broadcast()
//         }(i)
//     }

//     // c.L.Lock()
//     for ready != 10 {
//         c.Wait()
//         log.Println("裁判员被唤醒一次")
//     }
//     // c.L.Unlock()

//     //所有的运动员是否就绪
//     log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
// }

//使用 Cond 的另一个常见错误是，只调用了一次 Wait，没有检查等待条件是否满足，结果条件没满足，程序就继续执行了。
// func main() {
//     c := sync.NewCond(&sync.Mutex{})
//     var ready int

//     for i := 0; i < 10; i++ {
//         go func(i int) {
//             time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

//             // 加锁更改等待条件
//             c.L.Lock()
//             ready++
//             c.L.Unlock()

//             log.Printf("运动员#%d 已准备就绪\n", i)
//             // 广播唤醒所有的等待者
//             c.Broadcast()
//         }(i)
//     }

//     c.L.Lock()
//     // for ready != 10 {
//     c.Wait()
//     log.Println("裁判员被唤醒一次")
//     // }
//     c.L.Unlock()

//     //所有的运动员是否就绪
//     log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
// }
