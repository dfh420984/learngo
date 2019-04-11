package process2

import (
	"fmt"
	"learngo/chatroom/common/message"
	"encoding/json"
	"learngo/chatroom/server/utils"
	"net"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(mes *message.Message)  {
	//1.先取出SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &smsMes) err = ", err.Error())
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err = ", err.Error())
		return
	}
	//遍历在线用户开始群发
	for id, up := range usrMgr.onlineUsers {
		if id == smsMes.UserId { //过滤掉自己
			continue 
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn : conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendMesToEachOnlineUser tf.WritePkg(data) err", err.Error())
	}
}