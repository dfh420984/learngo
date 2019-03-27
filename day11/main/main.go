package main

import (
	"fmt"
	"learngo/day11/model"
	"sort"
	"math/rand"
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

//接口入门案例
func test06() {
	// phone := &model.Phone{}
	// camera := &model.Cammera{}
	// computer := &model.Computer{}
	// fmt.Printf("phone type %T \n", phone)
	// computer.Working(phone)
	// computer.Working(camera)
	//接口数组
	var usbArr [3]model.Usb
	usbArr[0] = model.Phone{"小米"}
	usbArr[1] = model.Phone{"vivo"}
	usbArr[2] = model.Cammera{"尼康"}
	var computer model.Computer
	for _,v := range usbArr { 
		computer.Working(v)
		fmt.Println()
	}
}

func test07() {
	var c1 model.C1
	c1.Say()
	c1.Hello()
	var a1 model.A1 = c1
	var b1 model.B1 = c1
	a1.Say()
	b1.Hello()
}

func test08() {
	var c1 model.C1
	var d1 model.D1 = c1
	d1.Say()
	d1.Hello()
	d1.Hi()
}

func test09() {
	f1 := &model.F1{}
	var e1 model.E1 = f1
	f1.Say()
	e1.Say()
	fmt.Printf("f1 type = %T \n", f1)
	fmt.Printf("e1 type = %T \n", e1)
}

//切片排序
func test10() {
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(s)
}

//结构体切片排序练习
func test11() {
	var heroes model.HeroSlice
	for i := 0; i < 10; i++ {
		hero := model.Hero{
			Name : fmt.Sprintf("英雄|%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println("排序前顺序")
	for _ , v := range heroes {
		fmt.Println(v)
	}

	sort.Sort(heroes)
	fmt.Println("排序后顺序")
	for _ , v := range heroes {
		fmt.Println(v)
	}
}

//学生结构体切片按分数排序练习
func test12() {
	var stues model.StuSlice
	for i := 0; i < 10; i++ {
		stu := model.Student{
			Name : fmt.Sprintf("学生%d", i + 1),
			Age : rand.Intn(20),
			Score : rand.Intn(100),
		}
		stues = append(stues, stu)
	} 
	fmt.Println("排序前顺序")
	for _ , v := range stues {
		fmt.Println(v)
	}

	sort.Sort(stues)
	fmt.Println("排序后顺序")
	for _ , v := range stues {
		fmt.Println(v)
	}
}

//继承和接口
func test13() {
	littleMokey := model.LittleMonkey{
		model.Monkey{
			Name : "孙悟空",
		},
	}
	littleMokey.Climb()
	littleMokey.Fly()
	littleMokey.Swim()
}

//接口类型断言
func test14() {
	var x interface{}
	var b float32 = 1.1
	x = b
	//y := x.(float32)
	if y, ok := x.(float32); ok {
		fmt.Printf("y得类型%T,值是%v", y, y)
	} else {
		fmt.Println("covert fail")
	}
	
}

func TypeJudge(items... interface{}) {
	for index, x := range items{
		switch x.(type) {
			case bool:
				fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
			case float32, float64:
				fmt.Printf("第%v个参数是float类型，值是%v\n", index, x)
			case int, int32, int64:
				fmt.Printf("第%v个参数是int类型，值是%v\n", index, x)
			case string:
				fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
			case model.Student:
				fmt.Printf("第%v个参数是Student类型，值是%v\n", index, x)
			case *model.Student:
				fmt.Printf("第%v个参数是*model.Student类型，值是%v\n", index, x)
			default:
				fmt.Printf("第%v个参数类型不确定，值是%v\n", index, x)
		}
			
	}
}

func main()  {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	//test09()
	//test10()
	//test11()
	//test12()
	//test13()
	//test14()
	var n1 float32 = 3.0
	var n2 float64 = 3.6
	var n3 int32 = 10
	var name string = "北京"
	address := "上海"
	n4 := 300 
	var stu  = model.Student{
		Name : "jack",
		Age : 30,
		Score : 99,
	}
	var stu2  = &model.Student{
		Name : "jack",
		Age : 30,
		Score : 99,
	}
	TypeJudge(n1,n2,n3,name,address,n4,stu,stu2)
}