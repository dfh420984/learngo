package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
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

//RPC golang gob 方式编码, 不能跨语言
func main() {
	//注册服务
	rpc.RegisterName("Calculate", new(Calculate))
	//绑定http
	rpc.HandleHTTP()
	//监听端口, 并启动服务
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
