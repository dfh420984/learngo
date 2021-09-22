package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	mes := make(chan int, 10)
	// producer
	for i := 0; i < 10; i++ {
		mes <- i
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// consumer
	go func(ctx context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-mes)
			}
		}
	}(ctx)

	defer close(mes)
	defer cancel()
	select {
	case <-ctx.Done():
		time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}

}
