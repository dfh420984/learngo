package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main()  { 
	fmt.Println("客户端准备开始连接")
	conn, err := net.Dial("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("客户端dial连接错误:%v \n", err)
		return 
	}
	line, err  := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("读取信息错误:%v", err)
	}
	n, err := conn.Write([]byte(line))
	if err != nil {
		fmt.Printf("write error:%v", err)
	}
	fmt.Printf("客户端发送了%d个字节，并退出\n" ,n)
}