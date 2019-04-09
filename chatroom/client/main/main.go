package main

import (
	"fmt"
	"os"
	"learngo/chatroom/client/process"
)

var (
	userId int
	userPwd string
	userName string
)

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
				fmt.Println("请输入用户的id")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("请输入用户的密码")
				fmt.Scanf("%s\n", &userPwd)
				// 完成登录
				//1. 创建一个UserProcess的实例
				up := &process.UserProcess{}
				up.Login(userId, userPwd)
			case 2:
				fmt.Println("注册用户")
				fmt.Println("请输入用户的id")
				fmt.Scanf("%d\n", &userId)
				fmt.Println("请输入用户的密码")
				fmt.Scanf("%s\n", &userPwd)
				fmt.Println("请输入用户名字")
				fmt.Scanf("%s\n", &userName)
				//1. 创建一个UserProcess的实例
				up := &process.UserProcess{}
				up.Register(userId, userPwd, userName)
			case 3:
				fmt.Println("退出系统")
				//loop = false
				os.Exit(0)
			default:
				fmt.Println("输入有误，请从新输入")
		}
	}
}