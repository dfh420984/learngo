package main

import (
	_ "fmt"
	"learngo/day12/myAccount/utils"
)

func main()  {
	myAccount := utils.FactoryAccount()
	myAccount.MainMenu()
	//fmt.Println(myAccount)
}