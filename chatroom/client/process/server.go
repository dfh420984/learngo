package process

import (
	"fmt"
	"net"
	"learngo/chatroom/client/utils"
	"os"
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
			fmt.Println("显示在线用户列表-")
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
		fmt.Printf("mes=%v \n", mes)
	}
}