package process2

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"learngo/chatroom/server/utils"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
}

//用来处理用户登陆逻辑
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	return tf.WritePkg(data)
}
