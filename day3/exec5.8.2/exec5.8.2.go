package main

import (
	"fmt"
)

func main()  {
	var name string
	var pwd string
	var num int = 3
	for i := 1; i <= num; i++ {
		fmt.Println("请输入姓名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name == "段福浩" && pwd == "888" {
			fmt.Println("恭喜登陆成功")
			break
		} else { 
			fmt.Printf("用户名或密码错误，你还有%v次机会", num - i)
		}
	}
}