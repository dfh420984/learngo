package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width  int
	Length int
}

func main() {
	//连接服务端
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	//请求服务
	//求面积
	arg1 := Params{Width: 10, Length: 20}
	res1 := 0
	err = client.Call("reactRpc.Area", arg1, &res1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("面积：%d\n", res1)
	//求周长
	res2 := 0
	err = client.Call("reactRpc.Perimeter", arg1, &res2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("周长：%d\n", res2)
}
