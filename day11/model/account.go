package model

import (
	"fmt"
)

type Account struct {
	AccountNo string
	Password string
	Balance float64
}

func (account *Account) Deposit(money float64, password string)  {
	if password != account.Password {
		fmt.Println("密码不正确")
		return 
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account) Draw(money float64, password string) {
	if password != account.Password {
		fmt.Println("密码不正确")
		return 
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account) Query(password string) {
	if password != account.Password {
		fmt.Println("密码不正确")
		return
	}
	fmt.Printf("balance=%v \n", account.Balance)
}


