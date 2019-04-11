package main

import (
	"fmt"
	"net"
	"learngo/chatroom/common/message"
	"learngo/chatroom/server/utils"
	"learngo/chatroom/server/process"
	"io"
)

type Processor struct {
	Conn net.Conn
}

//根据不同的消息类型来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) { 
	switch mes.Type {
		case message.LoginMesType:
			//处理登陆
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType:
			//处理注册
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType:
			//处理群发
			up := &process2.SmsProcess{}
			up.SendGroupMes(mes)
		case message.SendMesToOneType:
			//处理点对点聊天
			up := &process2.SmsProcess{}
			up.SendMesToOne(mes)
		case message.NotifyUserStatusMesType: 
			//处理用户退出登陆
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			up.DelOnlineUser(mes)
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) process2() (err error) {
	for { 
		//将要读取的数据包，封装程一个函数 返回Message,error
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出")
				return err
			} else {
				fmt.Println("server process readpkg err=", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}