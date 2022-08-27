package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calculate struct {
}

type Request struct {
	A int
	B int
}

type Response struct {
	Mul int
	Mod int
	Div int
}

//Multiple 乘法
func (c *Calculate) Multiple(req Request, res *Response) error {
	res.Mul = req.A * req.B
	return nil
}

//Divide 除法
func (c *Calculate) Divide(req Request, res *Response) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Div = req.A / req.B
	return nil
}

//取模
func (c *Calculate) Modulo(req Request, res *Response) error {
	if req.B == 0 {
		return errors.New("除数不能为0")
	}
	res.Mod = req.A % req.B
	return nil
}

//jsonrpc 支持跨语言
func main() {
	//注册服务
	rpc.RegisterName("Calculate", new(Calculate))
	//监听端口, 并启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	// 一直连接
	for {
		//接收数据
		con, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go func(con net.Conn) {
			fmt.Println("a new client")
			jsonrpc.ServeConn(con)
		}(con)
	}
}
