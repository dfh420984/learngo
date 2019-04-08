package process

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"learngo/chatroom/client/utils"
	"encoding/json"
	"encoding/binary"
	_"time"
)

type UserProcess struct {

}

func (this *UserProcess) Login(userId int, userPwd string)  (err error)  {
	// fmt.Printf("userId = %d userPwd = %s \n", userId, userPwd)
	// return nil
	//1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:8889")
	if err != nil {
		// handle error
		fmt.Println("dial error = ", err)
		return 
	}
	defer conn.Close()

	//2.先将消息结构化,然后在发送给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes Marshal err = ", err)
		return 
	}

	//5.将data付给mes.Data
	mes.Data = string(data)

	//6.将mes消息体实例序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes Marshal err = ", err)
		return 
	}

	//7.这个data就是我们要发送给服务器的数据
	//7.1先将data数据的长度发送给服务器，来防止数据丢包问题
	var pkglen uint32
	pkglen = uint32(len(data))
	//把获取到的字节长度写入byte切片中
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkglen)
	//开始发送数据
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("login conn.Write len err = ", err)
		return 
	}
	//fmt.Printf("客户端发送消息的长度=%d,内容=%s", len(data), string(data))

	//8.发送消息本身给服务器
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("login conn.Write data err = ", err)
		return 
	}

	// time.Sleep(time.Second*3)
	// fmt.Println("休眠了3秒")
	//9.处理服务器返回信息 
	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	mes, err = tf.ReadPkg() 
	if err != nil {
		fmt.Println("login readPkg(conn) error = ", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return 
}