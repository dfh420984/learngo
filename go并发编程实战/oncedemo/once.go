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

func netDial() {
	var addr = "baidu.com"
	once.Do(func() {
		conn, err = net.Dial("tcp", addr)
		fmt.Println(conn, err)
	})
}

// 值是3.0或者0.0的一个数据结构
var threeOnce struct {
	sync.Once
	v *Float
}

// 返回此数据结构的值，如果还没有初始化为3.0，则初始化
func three() *Float {
	threeOnce.Do(func() {
		// 使用Once初始化
		threeOnce.v = NewFloat(3.0)
	})
	return threeOnce.v
}
