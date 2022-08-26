package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Request struct {
	A int
	B int
}

type Response struct {
	Mul int
	Div int
	Mod int
}

func main() {
	cli, err := jsonrpc.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	var res Response
	req := &Request{
		A: 10,
		B: 5,
	}
	err = cli.Call("Calculate.Multiple", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("乘：", res.Mul)
	err = cli.Call("Calculate.Divide", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("除：", res.Div)
	err = cli.Call("Calculate.Modulo", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("模：", res.Mod)
}
