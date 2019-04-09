package process2

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"learngo/chatroom/server/utils"
	"learngo/chatroom/server/model"
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
	// if LoginMes.UserId == 100 && LoginMes.UserPwd == "123456" {
	// 	loginResMes.Code = 200
	// } else {
	// 	loginResMes.Code = 500
	// 	loginResMes.Error = "该用户不存在，请注册在使用"
	// }
	//去redis校验userId,userPwd
	user, err := model.MyUserDao.Login(LoginMes.UserId, LoginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
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

//用来处理用户注册逻辑
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mse.data中反序列取出RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("jsom Unmarshal fail err =", err)
	}

	//2.声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//3.在声明一个RegisterMes
	var registerResMes message.RegisterResMes

	//4.去redis完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发送未知错误。。"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println(registerMes.User, "注册成功")
	}

	//5.将registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json Marshal registerResMes fail err=", err)
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
