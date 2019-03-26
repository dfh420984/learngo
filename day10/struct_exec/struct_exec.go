package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Age int
	sex string
	Score [3]float64
	slice []int
	ptr *int
	map1 map[string]string
}

func (p Person) speak() {
	fmt.Println(p.Name,"是个goodman")
}

func (p Person) jisuan(n int) {
	res := 0
	for i := 1; i<=n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果=", res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

//struct声明定义赋值
func test() { 
	//1.第一种
	var person1 Person
	person1.Name = "段福浩"
	person1.Age = 31
	person1.sex = "男"
	fmt.Println(person1)

	//2.第二种（推荐）
	person2 := Person{"段福浩", 31, "男",[3]float64{60.0,70.0,80.0},make([]int,2),nil,make(map[string]string)}
	// person2.Score = [3]float64{60.0,70.0,80.0}
	// person2.slice = make([]int,2)
	// person2.ptr = nil
	// person2.map1 = make(map[string]string)
	fmt.Println(person2)

	//3.第三种
	var person3 *Person = new(Person)
	(*person3).Name = "段福浩"
	(*person3).Age = 31
	(*person3).sex = "男"
	fmt.Println(*person3)

	//4.第四种
	var person4 *Person = &Person{}
	(*person4).Name = "段福浩"
	(*person4).Age = 31
	(*person4).sex = "男"
	fmt.Println(*person4)
}

//结构体默认是值传递，如过要改变值需要使用&引用
func test02() {
	var p1 Person
	p1.Name = "段福浩"
	p1.Age = 31

	var p2 *Person = &p1
	(*p2).Age = 20

	fmt.Printf("p1.age=%v,(*p2).age=%v \n",p1.Age,(*p2).Age)
	fmt.Printf("p1得地址%p\n",&p1)
	fmt.Printf("p2得地址%p,p2值是%p\n",&p2,p2)
}

//结构体tag
type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

//结构体转json
func test03() {
	monster := Monster{"牛魔王",500,"芭蕉扇-"}
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}

//结构体绑定方法
type A struct {
	Num int
}

func (a A) test04() {
	fmt.Println(a.Num)
}

//课堂练习
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) area2() float64 {
	(*c).radius = 10.0
	return 3.14 * (*c).radius * (*c).radius
}

//实现String方法
type Student struct {
	Name string
	Gender string
	Age int
	Id int
	Score float64
}

func (stu Student) say() string {
	res := fmt.Sprintf("Name=%v,Gender=%v,Age=%v,Id=%v,Score=%v",stu.Name,stu.Gender,stu.Age,stu.Id,stu.Score)
	return res
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v],age=[%v]", stu.Name, stu.Age)
	return str
}

//课堂练习
type MethodUtils struct {
	//待定..
}

func (memthodUtils MethodUtils) printImage(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (memthodUtils MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func (memthodUtils MethodUtils) judegeNum(num int)  {
	if num % 2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

func (mu *MethodUtils) Print1(m int, n int, key string) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func (mu *MethodUtils) Print2() { 
	n := 0
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&n)
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Printf("%v * %v = %v ", j, i, j * i)
			}
			fmt.Println()
		}
	}
}


//结构体加减乘除练习
type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (cal *Calcuator) getRes(oper byte) float64 {
	res := 0.0
	switch oper {
		case '+':
			res = (*cal).Num1 + (*cal).Num2
		case '-':
			res = (*cal).Num1 - (*cal).Num2
		case '*':
			res = (*cal).Num1 * (*cal).Num2
		case '/':
			res = (*cal).Num1 / (*cal).Num2
	}
	return res
}

//结构体二维数组转换练习
type Exchange struct {
	Arr [3][3]int
}

func (ex Exchange) print3() {
	for i := 0; i < len(ex.Arr); i++ {
		for j := 0; j < len(ex.Arr[i]); j++ {
			fmt.Print(ex.Arr[j][i]," ")
		}
		fmt.Println()
	}

}

//求体积
type Box struct {
	len float64
	width float64
	height float64
}

func (b Box) tiji() float64 {
	return b.len * b.width * b.height
}


func main()  {
	//test()
	//test02()
	//test03()
	// a := A{123}
	// a.test04()
	// var p1 Person
	// p1.Name = "段福浩"
	// p1.speak()
	// p1.jisuan(10)
	// sum := p1.getSum(10,20)
	// fmt.Println("sum=",sum)
	// c := Circle{2.0}
	// res := c.area()
	// fmt.Println("res=",res)
	// res2 := (&c).area2()
	// fmt.Println("res2=",res2,"radius=",c.radius)
	// stu := Student{
	// 	"tom",
	// 	20,
	// }
	// fmt.Println(&stu)
	// var m MethodUtils
	// m.printImage(3,2)
	// area := m.area(5.0, 2.0)
	// fmt.Println("area=",area)
	// m.judegeNum(3)
	// (&m).Print1(3,2,"a")
	//  cal := &Calcuator{Num1:10.0, Num2:2.0}
	//  res := cal.getRes('+')
	//  fmt.Println("res=", res)

	// var mu *MethodUtils = &MethodUtils{}
	// mu.Print2()
	// ex := Exchange{
	// 	[3][3]int{{1,2,3},{4,5,6},{7,8,9}},
	// }
	// ex.print3()

	// stu := Student{
	// 	Name:"段福浩",
	// 	Gender:"男",
	// 	Age:30,
	// 	Id:123,
	// 	Score:100.0,
	// }
	// fmt.Println(stu.say())

	var b Box
	fmt.Println("请输入长度")
	fmt.Scanln(&b.len)
	fmt.Println("请输入宽度")
	fmt.Scanln(&b.width)
	fmt.Println("请输入高度")
	fmt.Scanln(&b.height)
	fmt.Println(b.tiji())
}