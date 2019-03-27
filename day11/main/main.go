package main

import (
	"fmt"
	"learngo/day11/model"
)

func test01() {
	account := model.Account{
		AccountNo : "1234",
		Password : "888888",
		Balance : 10000.0,
	}
	account.Deposit(10000.0, "888888")
	account.Query("888888")
	account.Draw(5000.0, "888888")
	account.Query("888888")
}

func test02() {
	per := model.Newperson("jack")
	per.SetAge(30)
	per.SetSal(6000)
	fmt.Println("age=", per.GetAge())
	fmt.Println("sal=", per.GetSal())
}

func test03() {
	p1 := &model.Puiple{}
	p1.Student.Name = "tom"
	p1.Student.Age = 10
	p1.Student.SetScore(90)
	p1.Testing()
	p1.Student.ShowInfo()

	p2 := &model.Graduate{}
	p2.Student.Name = "jack"
	p2.Student.Age = 23
	p2.Student.SetScore(120)
	p2.Testing()
	p2.Student.ShowInfo()
}

func test04() {
	c := model.C{
		model.A{
			Name : "jack",
			Age : 20,
		},
		model.B{
			Name : "Tom",
			Score : 99.0,
		},
	}
	fmt.Println("c=", c)
	fmt.Println("c.A.Name=", c.A.Name)
}

func test05() {
	d := model.D{
		&model.A{
			Name : "jack",
			Age : 20,
		},
		&model.B{
			Name : "Tom",
			Score : 99.0,
		},
	}
	fmt.Println("d=", d)
	fmt.Printf("*d.A type=%T \n", *d.A)
	fmt.Printf("d.B type=%T \n", d.B)
	fmt.Println(d.B.Score)
}

func main()  {
	//test01()
	//test02()
	//test03()
	//test04()
	test05()
}