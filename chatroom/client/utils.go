package main

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

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