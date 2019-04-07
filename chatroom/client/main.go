package main

import (
	"fmt"
	"os"
)

var userId int
var userPwd string

func main()  { 
	var key int
	loop := true
	for loop {
		fmt.Println("------------------欢迎登录多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登陆聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统") 
		fmt.Println("请选择1-3:")
		fmt.Scanf("%d\n", &key)
		switch key {
			case 1:
				fmt.Println("登陆聊天室")
				loop = false
			case 2:
				fmt.Println("注册用户")
				loop = false
			case 3:
				fmt.Println("退出系统")
				//loop = false
				os.Exit(0)
			default:
				fmt.Println("输入有误，请从新输入")
		}
	}
	if key == 1 {
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)
		err := login(userId, userPwd)
		if err != nil {
			fmt.Println("登陆失败")
		} else {
			fmt.Println("登陆成功")
		}
	} else if key == 2 {
		fmt.Println("注册用户逻辑")
	}
}