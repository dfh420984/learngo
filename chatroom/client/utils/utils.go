package utils

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"encoding/json"
	"encoding/binary"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取服务端发送的数据。。。")
	//conn在没有关闭的情况下Read会阻塞，关闭了就不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("conn.read() err = ", err)
		return 
	}
	//fmt.Println("读取到的buf=", buf[:4])
	//将buf[:4]转换成uint32类型
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		return 
	}
	err = json.Unmarshal(this.Buf[:pkglen], &mes) 
	if err != nil {
		fmt.Println("server json.Unmarshal  err = ", err)
		return 
	}
	return 
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return 
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return 
	}
	return 
}