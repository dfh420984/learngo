package process

import (
	"fmt"
	"learngo/chatroom/common/message"
	"learngo/chatroom/client/model"
	"encoding/json"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //我们在用户登录成功后，完成对CurUser初始化

//在客户端显示在线用户列表
func outputOnlineUser()  {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}

//处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status 
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}

//输出群发消息
func outputGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal err=", err.Error())
		return
	}
	info := fmt.Sprintf("用户id:\t%d,对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}

//输出点对点消息
func outputToOneMes(mes *message.Message) {
	var sendMesToOne message.SendMesToOne
	err := json.Unmarshal([]byte(mes.Data), &sendMesToOne)
	if err != nil {
		fmt.Println("outputToOneMes json.Unmarshal err=", err.Error())
		return
	}
	info := fmt.Sprintf("接收到用户id:\t%d,对id:%d 用户说:\t%s", sendMesToOne.UserId, sendMesToOne.ReciverId,
	sendMesToOne.Content)
	fmt.Println(info)
}