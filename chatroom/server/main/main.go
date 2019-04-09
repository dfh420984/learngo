package main

import (
	"fmt"
	"net"
	"learngo/chatroom/server/model"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main()  {
	//初始化redis链接池
	initPool("0.0.0.0:6379", 16, 0, 300 * time.Second)
	initUserDao()
	fmt.Println("服务器正在监听8889端口")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen err = ", err)
		return 
	}
	//服务器等待客户端连接，发送数据
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err = ", err)
		}
		//如果连接成功，开启一个协程于客户端通讯
		go process(conn)
	}
	
}

func process(conn net.Conn) {
	defer conn.Close()
	process := &Processor{
		Conn : conn,
	}
	err := process.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯错误err=", err)
		return
	}
}