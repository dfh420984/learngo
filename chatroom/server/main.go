package main

import (
	"fmt"
	"net"
	"io"
	"learngo/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

func main()  {
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
	for { 
		//将要读取的数据包，封装程一个函数 返回Message,error
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return 
			} else {
				fmt.Println("server process readpkg err=", err)
			}
		}
		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

//根据不同的消息类型来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
		case message.LoginMesType:
			//处理登陆
			err = serverProcessLogin(conn, mes)
		case message.RegisterMesType:
			//处理注册
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据。。。")
	//conn在没有关闭的情况下Read会阻塞，关闭了就不会阻塞
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn.read() err = ", err)
		return 
	}
	//fmt.Println("读取到的buf=", buf[:4])
	//将buf[:4]转换成uint32类型
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(buf[0:4])
	n, err := conn.Read(buf[:pkglen])
	if n != int(pkglen) || err != nil {
		return 
	}
	err = json.Unmarshal(buf[:pkglen], &mes) 
	if err != nil {
		fmt.Println("server json.Unmarshal  err = ", err)
		return 
	}
	return 
}

func writePkg(conn net.Conn, data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data)) 
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return 
	}

	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return 
	}
	return 
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.先从mse.data中反序列取出LoginMes
	var LoginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &LoginMes)
	if err != nil {
		fmt.Println("jsom Unmarshal fail err =", err)
	}

	//2.声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//3.在声明一个loginResMes
	var loginResMes message.LoginResMes
	if LoginMes.UserId == 100 && LoginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册在使用"
	}
	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json Marshal loginResMes fail err=", err)
		return
	}
	resMes.Data = string(data)
	//4.将resMes消息体实例序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("resMes Marshal err = ", err)
		return 
	}

	//5.将结果返回给客户端
	return writePkg(conn, data)
}