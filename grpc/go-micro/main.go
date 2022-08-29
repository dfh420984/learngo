package main

import (
	"context"
	"fmt"

	"dfhgrpc.168.cn/go-micro/proto"
	"github.com/micro/micro/v3/service"
)

// 声明结构体
type Hello struct {
}

// 实现接口方法
func (h *Hello) Info(ctx context.Context, req *proto.InfoRequest, resp *proto.UserResponse) (err error) {
	resp.Msg = "你好" + req.Username
	return
}

// micro server
// cd go-micro
// micro run .
// micro go-micro hello info --username=zs
// micro go-micro --help
func main() {
	//1. 得到微服务实例
	ser := service.New(
		//设置微服务名字,用来做访问用的
		service.Name("hello"),
	)
	//2. 初始化
	ser.Init()
	//3. 服务注册
	err := proto.RegisterHelloHandler(ser.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	//4. 启动服务
	err = ser.Run()
	if err != nil {
		fmt.Println(err)
	}
}
