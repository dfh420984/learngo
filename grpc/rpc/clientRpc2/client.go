package main

import (
	"fmt"
	"log"
	"net/rpc"
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
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	var res Response
	req := &Request{
		A: 10,
		B: 5,
	}
	err = cli.Call("Calculate.Multiple", req, &res)
	err = cli.Call("Calculate.Divide", req, &res)
	err = cli.Call("Calculate.Modulo", req, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
