package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//rpc 结构体，参数，函数必须都是首字母大写
type React struct {
}

type Params struct {
	Width  int
	Length int
}

//Area rpc 方法第一个参数是接收参数，第二个参数是返回给客户端值, 必须是指针
func (r *React) Area(p Params, res *int) error {
	*res = p.Width * p.Length
	return nil
}

//Perimeter 周长
func (r *React) Perimeter(p Params, res *int) error {
	*res = (p.Width + p.Length) * 2
	return nil
}

//注册服务对象，启动服务
func main() {
	// 注册服务
	react := new(React)
	err := rpc.RegisterName("reactRpc", react)
	if err != nil {
		log.Fatal(err)
	}
	// 通过HandleHTTP将React提供服务绑定到http协议上，为调用者进行http协议传输数据
	rpc.HandleHTTP()
	//指定监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	//开启服务
	http.Serve(listen, nil)
}
