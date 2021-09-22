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

//常见问题一：计数器设置为负值
//waitgroup常见错误，WaitGroup 的计数器的值必须大于等于 0。我们在更改这个计数值的时候，WaitGroup 会先做检查，如果计数值被设置为负数，就会导致 panic。

//计数器的初始值为 10，当第一次传入 -10 的时候，计数值被设置为 0，不会有啥问题。但是，再紧接着传入 -1 以后，计数值就被设置为负数了，程序就会出现 panic。
// func main() {
//     var wg sync.WaitGroup
//     wg.Add(10)

//     wg.Add(-10)//将-10作为参数调用Add，计数值被设置为0

//     wg.Add(-1)//将-1作为参数调用Add，如果加上-1计数值就会变为负数。这是不对的，所以会触发panic
// }

//多调用了一次 Done 方法后，会导致计数值为负，所以程序运行到这一行会出现 panic。
// func main() {
//     var wg sync.WaitGroup
//     wg.Add(1)

//     wg.Done()

//     wg.Done()
// }

//常见问题二：不期望的 Add 时机
//在这个例子中，我们原本设想的是，等四个 goroutine 都执行完毕后输出 Done 的信息，但是它的错误之处在于，将 WaitGroup.Add 方法的调用放在了子 gorotuine 中。等主 goorutine 调用 Wait 的时候，因为四个任务 goroutine 一开始都休眠，所以可能 WaitGroup 的 Add 方法还没有被调用，WaitGroup 的计数还是 0，所以它并没有等待四个子 goroutine 执行完毕才继续执行，而是立刻执行了下一步。
// func main() {
//     var wg sync.WaitGroup
//     go dosomething(100, &wg) // 启动第一个goroutine
//     go dosomething(110, &wg) // 启动第二个goroutine
//     go dosomething(120, &wg) // 启动第三个goroutine
//     go dosomething(130, &wg) // 启动第四个goroutine

//     wg.Wait() // 主goroutine等待完成
//     fmt.Println("Done")
// }

// func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
//     duration := millisecs * time.Millisecond
//     time.Sleep(duration) // 故意sleep一段时间

//     wg.Add(1)
//     fmt.Println("后台执行, duration:", duration)
//     wg.Done()
// }

//导致这个错误的原因是，没有遵循先完成所有的 Add 之后才 Wait。要解决这个问题，一个方法是，预先设置计数值：
// func main() {
//     var wg sync.WaitGroup
//     wg.Add(4) // 预先设定WaitGroup的计数值

//     go dosomething(100, &wg) // 启动第一个goroutine
//     go dosomething(110, &wg) // 启动第二个goroutine
//     go dosomething(120, &wg) // 启动第三个goroutine
//     go dosomething(130, &wg) // 启动第四个goroutine

//     wg.Wait() // 主goroutine等待
//     fmt.Println("Done")
// }

// func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
//     duration := millisecs * time.Millisecond
//     time.Sleep(duration)

//     fmt.Println("后台执行, duration:", duration)
//     wg.Done()
// }

//另一种方法是在启动子 goroutine 之前才调用 Add：（推荐这样使用，避免增加携程调用时，忘记修改Add数量）
// func main() {
//     var wg sync.WaitGroup

//     dosomething(100, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//     dosomething(110, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//     dosomething(120, &wg) // 调用方法，把计数值加1，并启动任务goroutine
//     dosomething(130, &wg) // 调用方法，把计数值加1，并启动任务goroutine

//     wg.Wait() // 主goroutine等待，代码逻辑保证了四次Add(1)都已经执行完了
//     fmt.Println("Done")
// }

// func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
//     wg.Add(1) // 计数值加1，再启动goroutine

//     go func() {
//         duration := millisecs * time.Millisecond
//         time.Sleep(duration)
//         fmt.Println("后台执行, duration:", duration)
//         wg.Done()
//     }()
// }


//常见问题三：前一个 Wait 还没结束就重用 WaitGroup
// 如果我们在 WaitGroup 的计数值还没有恢复到零值的时候就重用，就会导致程序 panic。我们看一个例子，初始设置 WaitGroup 的计数值为 1，启动一个 goroutine 先调用 Done 方法，接着就调用 Add 方法，Add 方法有可能和主 goroutine 并发执行。
// func main() {
//     var wg sync.WaitGroup
//     wg.Add(1)
//     go func() {
//         time.Sleep(time.Millisecond)
//         wg.Done() // 计数器减1
//         wg.Add(1) // 计数值加1
//     }()
//     wg.Wait() // 主goroutine等待，有可能和第7行并发执行
// }