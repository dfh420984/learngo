package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

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
