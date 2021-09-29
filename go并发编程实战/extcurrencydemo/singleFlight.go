package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	sf           = singleflight.Group{}
	requestCount = int64(0)
	resp         = make(chan int64)
	wg           sync.WaitGroup
)

// Do：这个方法执行一个函数，并返回函数执行的结果。你需要提供一个 key，对于同一个 key，在同一时间只有一个在执行，同一个 key 并发的请求会等待。第一个执行的请求返回的结果，就是它的返回结果。函数 fn 是一个无参的函数，返回一个结果或者 error，而 Do 方法会返回函数执行的结果或者是 error，shared 会指示 v 是否返回给多个请求。
// DoChan：类似 Do 方法，只不过是返回一个 chan，等 fn 函数执行完，产生了结果以后，就能从这个 chan 中接收这个结果。
// Forget：告诉 Group 忘记这个 key。这样一来，之后这个 key 请求会执行 f，而不是等待前一个未完成的 fn 函数的结果。

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			do, err, _ := sf.Do("number", Request)
			if err != nil {
				log.Println(err)
			}
			log.Println("resp", do)
		}()
	}
	time.Sleep(1 * time.Second)
	resp <- atomic.LoadInt64(&requestCount)
	wg.Wait()
}

func Request() (interface{}, error) {
	atomic.AddInt64(&requestCount, 1)
	return <-resp, nil
}
