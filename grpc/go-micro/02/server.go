package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "dfhgrpc.168.cn/go-micro/02/proto"
	"github.com/micro/micro/v3/service"
)

type Example struct {
}

type Foo struct {
}

func (e *Example) Call(ctx context.Context, req *pb.CallRequest, resp *pb.CallResponse) error {
	log.Println("收到Example.Call请求")
	if len(req.Name) == 0 {
		return errors.New("go.micro.api.example no context")
	}
	resp.Message = "RPC Example.Call 收到请求：" + req.Name
	return nil
}

func (f *Foo) Bar(ctx context.Context, req *pb.EmptyRequest, resp *pb.EmptyResponse) error {
	log.Println("收到 Foo.Bar请求")
	return nil
}

func main() {
	//1.得到微服务实例
	ser := service.New(
		service.Name("go.micro.api.example"),
	)
	//2.初始化
	ser.Init()
	//3.注册example接口
	err := pb.RegisterExampleHandler(ser.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}
	//4.注册foo接口
	err = pb.RegisterFooHandler(ser.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	//4. 启动服务
	err = ser.Run()
	if err != nil {
		fmt.Println(err)
	}
}
