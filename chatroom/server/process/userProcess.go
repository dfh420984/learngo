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
	//服务器连接
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户
	UserId int
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
		//将登陆成功的用户放入usrMgr实例中
		this.UserId = LoginMes.UserId
		usrMgr.AddOnlineUser(this)
		//将userId放入loginResMes中
		for id, _ := range usrMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		//通知其他用户上线
		this.NotifyOthersOnlineUser(LoginMes.UserId,message.UserOnline)
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

//通知所有在线用户的方法
func (this *UserProcess) NotifyOthersOnlineUser(userId int, status int) {
	for id, up := range usrMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始一个个通知
		up.NotifyMeOnline(userId, status)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int, status int) { 
	//1.开始组装NotifyUserStatusMes结构体消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = status

	//2.将NotifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err =", err)
		return
	}

	//3.对mes进行序列化准备发送
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err =", err)
		return
	}

	//4.创建Transfer实例发送
	tf := &utils.Transfer{
		Conn : this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyOnMe err", err)
	}
	return
}

func (this *UserProcess) DelOnlineUser(mes *message.Message) {
	var notifyUserStatusMes message.NotifyUserStatusMes
	err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("UpdateUserStatus json.Unmarshal error =", err)
		return
	}
	//更新onlineUsers状态
	usrMgr.DelOnlineUser(notifyUserStatusMes.UserId)
	//通知所有其他用户
	this.NotifyOthersOnlineUser(notifyUserStatusMes.UserId,message.UserOffline)
}
