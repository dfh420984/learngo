package main

import (
	"fmt"
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

func main()  {
	//test()
	test02()
}