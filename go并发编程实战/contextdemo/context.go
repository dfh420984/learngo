package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

//通过channel控制
// func main() {
// 	messages := make(chan int, 10)
// 	done := make(chan bool)
// 	defer close(messages)
// 	//consumer
// 	go func() {
// 		ticker := time.NewTicker(1 * time.Second)
// 		for _ = range ticker.C {
// 			select {
// 			case <-done:
// 				fmt.Println("child process interrupt...")
// 				return
// 			default:
// 				fmt.Printf("send message: %d\n", <-messages)
// 			}
// 		}
// 	}()
// 	// producer
// 	for i := 1; i < 10; i++ {
// 		messages <- i
// 	}
// 	time.Sleep(5 * time.Second)
// 	close(done)
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("main process exit!")
// }

// 考虑下面这种情况：假如主协程中有多个任务1, 2, …m，主协程对这些任务有超时控制；而其中任务1又有多个子任务1, 2, …n，任务1对这些子任务也有自己的超时控制，那么这些子任务既要感知主协程的取消信号，也需要感知任务1的取消信号。

// 如果还是使用done channel的用法，我们需要定义两个done channel，子任务们需要同时监听这两个done channel。嗯，这样其实好像也还行哈。但是如果层级更深，如果这些子任务还有子任务，那么使用done channel的方式将会变得非常繁琐且混乱。

// 我们需要一种优雅的方案来实现这样一种机制：

// 上层任务取消后，所有的下层任务都会被取消；中间某一层的任务取消后，只会将当前任务的下层任务取消，而不会影响上层的任务以及同级任务。
//WithValue demo 链式查找
// func main() {
// 	ctx := context.TODO()
// 	ctx = context.WithValue(ctx, "key1", "0001")
// 	ctx = context.WithValue(ctx, "key2", "0001")
// 	ctx = context.WithValue(ctx, "key3", "0001")
// 	ctx = context.WithValue(ctx, "key4", "0004")

// 	fmt.Println(ctx.Value("key1"))
// }

//WithCancel demo
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go func() {
// 		defer func() {
// 			fmt.Println("goroutine exit")
// 		}()
// 		for {
// 			select {
// 			case val := <-ctx.Done():
// 				fmt.Println(val)
// 				return
// 			default:
// 				time.Sleep(time.Second)
// 			}
// 		}
// 	}()
// 	time.Sleep(time.Second)
// 	cancel()
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("main end")
// }

//WithTimeout demo
// func main() {
// 	//设置1秒超时
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()
// 	//开启一个500ms任务
// 	//go handel(ctx, 500*time.Millisecond)
// 	go handel(ctx, 1500*time.Millisecond)
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("main end")
// }

// func handel(ctx context.Context, d time.Duration) {
// 	select {
// 	case <-ctx.Done(): //ctx超时关闭时进入
// 		fmt.Println("handle", ctx.Err())
// 	case <-time.After(d): //模拟处理耗时500ms任务
// 		fmt.Println("process request with", d)
// 	}
// }

//WithCancel demo
func main() {
	parent := context.Background()                              //root parent
	cancelCtx, cancel := context.WithCancel(parent)             //parent
	valueCtx := context.WithValue(cancelCtx, "test", "context") //child
	go func() {                                                 //valueCtx
		for {
			select {
			case <-valueCtx.Done():
				log.Printf("valueCtx done")
				return
			default:
				res := valueCtx.Value("test")
				log.Printf("valueCtx working:" + res.(string))
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			select {
			case <-cancelCtx.Done():
				log.Printf("cancelCtx game over")
				return
			default:
				log.Printf("cancelCtx working:")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("main end")

}
