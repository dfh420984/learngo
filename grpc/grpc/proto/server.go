package main

import (
	"fmt"
	"net"

	"dfhgrpc.168.cn/grpc/proto/user"
	grpc "google.golang.org/grpc"
)

//程序执行入口
func main() {
	//1.监听
	addr := "127.0.0.1:8080"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常 err:%v\n", err)
	}
	fmt.Printf("开始监听：%s\n", addr)
	//2.实例化grpc
	s := grpc.NewServer()
	//3.在grpc上注册服务
	user.RegisterUserInfoServiceServer(s, user.NewUserInfoService())
	//4.启动grpc服务端
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("server启动异常 err:%v\n", err)
	}
}
