package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n , err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err = %v \n", err)
			return 
		}
		fmt.Println(string(buf[:n]))
	}
}

func main()  {
	fmt.Println("服务器开始监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("服务器listen连接错误:%v \n", err)
	}
	defer listen.Close()
	//循环等待客户端连接
	for {
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() error=", err)
		} else {
			fmt.Printf("listen.Accept() success conn = %v, 客户端ip = %v \n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}