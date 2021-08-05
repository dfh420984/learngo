package main

import (
	"fmt"
	"net"

	"dfhgrpc.168.cn/ch13/13.3/grpc/protocol"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:18887"
)

func main() {
	//监听网络
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("网络异常", err)
	}
	//创建grpc句柄
	srv := grpc.NewServer()
	//将server结构体注册到 grpc服务中
}
