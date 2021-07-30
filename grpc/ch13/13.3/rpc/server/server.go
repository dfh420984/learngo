package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line), *ack)
	return nil
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:13133")
	if err != nil {
		log.Fatal("err")
	}
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务端启动")
	listener := new(Listener)
	rpc.Register(listener)
	rpc.Accept(lis)
}
