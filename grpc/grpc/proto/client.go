package main

import (
	"context"
	"fmt"

	"dfhgrpc.168.cn/grpc/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var addr = "127.0.0.1:8080"
	//1.创建grpc服务端链接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("链接异常", err)
	}
	defer conn.Close()
	//实力化grpc客户端
	client := user.NewUserInfoServiceClient(conn)
	//组装参数
	req := user.UserRequest{
		Name: "zs",
	}
	resp, err := client.GetUserInfo(context.Background(), &req)
	if err != nil {
		fmt.Println("响应异常", err)
	}
	fmt.Println("响应结果", resp)
}
