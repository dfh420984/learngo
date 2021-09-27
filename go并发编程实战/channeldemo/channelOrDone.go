package main

import (
	"fmt"
	"reflect"
	"time"
)

// 首先来看 Or-Done 模式。Or-Done 模式是信号通知模式中更宽泛的一种模式。这里提到了“信号通知模式”，我先来解释一下。我们会使用“信号通知”实现某个任务执行完成后的通知机制，在实现时，我们为这个任务定义一个类型为 chan struct{}类型的 done 变量，等任务结束后，我们就可以 close 这个变量，然后，其它 receiver 就会收到这个通知。这是有一个任务的情况，如果有多个任务，只要有任意一个任务执行完，我们就想获得这个信号，这就是 Or-Done 模式。比如，你发送同一个请求到多个微服务节点，只要任意一个微服务节点返回结果，就算成功，这个时候，就可以参考下面的实现：

func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[1]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// switch len(channels) {
		// case 2:
		// 	select {
		// 	case <-channels[0]:
		// 	case <-channels[1]:
		// 	}
		// default:
		// 	select {
		// 	case <-channels[0]:
		// 	case <-channels[1]:
		// 	case <-channels[2]:
		// 	case <-or(append(channels[3:], orDone)...):
		// 	}
		// }

		//利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		fmt.Println(reflect.Select(cases))
	}()
	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-or(sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}
