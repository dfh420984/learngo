package main

import (
	"fmt"
	"net"
	"sync"
)

var (
	once sync.Once
	conn net.Conn
	err  error
)

//很多时候我们是要延迟进行初始化的，所以有时候单例资源的初始化，我们会使用下面的方法：
//这种方式虽然实现起来简单，但是有性能问题。一旦连接创建好，每次请求的时候还是得竞争锁才能读取到这个连接，这是比较浪费资源的，因为连接如果创建好之后，其实就不需要锁的保护了
// 使用互斥锁保证线程(goroutine)安全
// var connMu sync.Mutex
// var conn net.Conn

// func getConn() net.Conn {
//     connMu.Lock()
//     defer connMu.Unlock()

//     // 返回已创建好的连接
//     if conn != nil {
//         return conn
//     }

//     // 创建连接
//     conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
//     return conn
// }

// // 使用连接
// func main() {
//     conn := getConn()
//     if conn == nil {
//         panic("conn is nil")
//     }
// }

//sync.Once 只暴露了一个方法 Do，你可以多次调用 Do 方法，但是只有第一次调用 Do 方法时 f 参数才会执行，这里的 f 是一个无参数无返回值的函数。
//func (o *Once) Do(f func())

func main() {
	//第一个函数
	f1 := func() {
		fmt.Println("in f1")
	}
	once.Do(f1)
	//第二个函数
	f2 := func() {
		fmt.Println("in f2")
	}
	once.Do(f2)
	fmt.Println("main end")
}

//f 参数是一个无参数无返回的函数，所以你可能会通过闭包的方式引用外面的参数，比如
// func main() {
//     conn := getConnOnce()
//     if conn == nil {
//         panic("conn is nil")
//     }
// }

// func getConnOnce() net.Conn {
// 	connMu.Lock()
// 	defer connMu.Unlock()

// 	// 返回已创建好的连接
// 	if conn != nil {
// 		return conn
// 	}

// 	// 创建连接
// 	var addr = "baidu.com"
// 	var conn net.Conn
// 	once.Do(func() {
// 		conn, _ = net.Dial("tcp", addr)
// 	})
// 	return conn
// }


// 使用 Once 可能出现的 2 种错误
// 第一种错误：死锁
// func main() {
//     var once sync.Once
//     once.Do(func() {
//         once.Do(func() {
//             fmt.Println("初始化")
//         })
//     })
// }


//比如下面的例子，由于一些防火墙的原因，googleConn 并没有被正确的初始化，后面如果想当然认为既然执行了 Do 方法 googleConn 就已经初始化的话，会抛出空指针的错误：
// func main() {
//     var once sync.Once
//     var googleConn net.Conn // 到Google网站的一个连接

//     once.Do(func() {
//         // 建立到google.com的连接，有可能因为网络的原因，googleConn并没有建立成功，此时它的值为nil
//         googleConn, _ = net.Dial("tcp", "google.com:80")
//     })
//     // 发送http请求
//     googleConn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
//     io.Copy(os.Stdout, googleConn)
// }
