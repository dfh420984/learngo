package utils

import (
	"fmt"
)

type MyAccount struct {
	key string
	balance float64
	money float64
	details string
	note string
	loop bool
	flag bool
}

func FactoryAccount() *MyAccount {
	return &MyAccount{
		key : "",
		balance : 10000.0,
		money : 0.0,
		details : "收支\t账户金额\t收支金额\t说    明",
		note : "",
		loop : true,
		flag : false,
	}
}

func (this *MyAccount) showDetails() { 
	fmt.Println("-----------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("还没有收支详情，快去操作吧！")
	}

}

func (this *MyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	//收入    11000           1000            有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *MyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		return
	}
	this.balance -= this.money // 修改账户余额
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *MyAccount) exit() {
	fmt.Println("你确定要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}

	if choice == "y" {
		this.loop = false
	}
}

func (this *MyAccount)  MainMenu() { 
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4)：")
		fmt.Scanln(&this.key)
		switch this.key {
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default:
				fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}
	}
	
}