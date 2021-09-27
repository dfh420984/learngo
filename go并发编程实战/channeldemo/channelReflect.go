package main

import (
	"fmt"
	"reflect"
)

// select 语句可以处理 chan 的 send 和 recv，send 和 recv 都可以作为 case clause。如果我们同时处理两个 chan，就可以写成下面的样子：
// select {
// case v := <-ch1:
// 	fmt.Println(v)
// case v := <-ch2:
// 	fmt.Println(v)
// }

// 如果需要处理三个 chan，你就可以再添加一个 case clause，用它来处理第三个 chan。可是，如果要处理 100 个 chan 呢？一万个 chan 呢？
// chan 的数量在编译的时候是不定的，在运行的时候需要处理一个 slice of chan，这个时候，也没有办法在编译前写成字面意义的 select。那
//使用反射操作 Channel
// 下面，我来借助一个例子，来演示一下，动态处理两个 chan 的情形。因为这样的方式可以动态处理 case 数据，所以，你可以传入几百几千几万的 chan，这就解决了不能动态处理 n 个 chan 的问题。首先，createCases 函数分别为每个 chan 生成了 recv case 和 send case，并返回一个 reflect.SelectCase 数组。然后，通过一个循环 10 次的 for 循环执行 reflect.Select，这个方法会从 cases 中选择一个 case 执行。第一次肯定是 send case，因为此时 chan 还没有元素，recv 还不可用。等 chan 中有了数据以后，recv case 就可以被选择了。这样，你就可以处理不定数量的 chan 了。

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	// 创建SelectCase,createCases 函数分别为每个 chan 生成了 recv case 和 send case，并返回一个 reflect.SelectCase 数组
	var cases = createCases(ch1, ch2)
	// 执行10次select
	for i := 0; i < 10; i++ {
		//通过 reflect.Select 函数，你可以将一组运行时的 case clause 传入，当作参数执行。Go 的 select 是伪随机的，它可以在执行的 case 中随机选择一个 case，并把选择的这个 case 的索引（chosen）返回，如果没有可用的 case 返回，会返回一个 bool 类型的返回值，这个返回值用来表示是否有 recv case 成功被选择。如果是 recv case，还会返回接收的元素
		chosen, recv, recvOK := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", chosen, cases[chosen].Dir, recv, recvOK)
		} else { //send case
			fmt.Println("send:", chosen, cases[chosen].Dir, recvOK)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase
	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	// 创建send case
	for i, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir: reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(i),
		})
	}
	return cases
}