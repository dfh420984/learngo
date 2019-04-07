package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println("服务器正在监听8889端口")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err = ", err)
		return 
	}
	//服务器等待客户端连接，发送数据
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
		}
		//如果连接成功，开启一个协程于客户端通讯
		go process(conn)
	}
	
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 8096)
		fmt.Println("读取客户端发送的数据。。。")
		n, err := conn.Read(buf[:4])
		if n != 4 || err !=nil {
			fmt.Println("conn.read() err = ", err)
			return 
		}
		fmt.Println("读取到的buf=", buf[:4])
	}
}