package process

import (
	"fmt"
	"net"
	"learngo/chatroom/client/utils"
	"learngo/chatroom/common/message"
	"os"
	"encoding/json"
)

func ShowMenu()  {
	fmt.Println("-------恭喜xxx登录成功---------")
	fmt.Println("-------1. 显示在线用户列表---------")
	fmt.Println("-------2. 发送消息---------")
	fmt.Println("-------3. 信息列表---------")
	fmt.Println("-------4. 退出系统---------")
	fmt.Println("请选择(1-4):")
	var key int 
	fmt.Scanf("%d\n", &key) 
	switch key {
		case 1:
			fmt.Println("在线用户列表:")
			outputOnlineUser()
		case 2:
			fmt.Println("发送消息)")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出了系统...")
			os.Exit(0)
		default :
			fmt.Println("你输入的选项不正确..")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn : conn,
	}
	for { 
		fmt.Println("客户端正在等待服务器端发送数据")
		mes, err := tf.ReadPkg() 
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}
		//读取到进行下一步处理逻辑
		//fmt.Printf("mes=%v \n", mes)
		switch mes.Type {
			case message.NotifyUserStatusMesType: //有人上线了
				var notifyUserStatusMes message.NotifyUserStatusMes 
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes) 
				//将用户信息保存到onlineUsers map[int]*message.User
				updateUserStatus(&notifyUserStatusMes)
			default:
				fmt.Println("服务器返回了未知信息")
		}
	}
}