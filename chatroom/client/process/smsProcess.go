package process

import (
	"fmt"
	"learngo/chatroom/common/message"
	"encoding/json"
	"learngo/chatroom/client/utils"
)

type SmsProcess struct {

}

func (this *SmsProcess) SendGroupMes(content string) (err error)  {
	//1.创建mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.声明SmsMes
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//3.序列化SmsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err = ", err)
		return 
	}

	//4.在序列化mes
	mes.Data = string(data)
	data, err = json.Marshal(mes)

	//5.给服务器发送消息
	tf := &utils.Transfer{
		Conn : CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err=", err.Error())
		return
	}
	return 
}