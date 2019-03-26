package model

import (
	"fmt"
)

type person struct {
	Name string
	age int
	sal float64
}

func Newperson(name string) *person {
	return &person{
		Name : name,
	}
}

func (per *person) SetAge(age int) { 
	if age < 0 || age > 150 {
		fmt.Println("年龄范围不正确")
	}
	per.age = age
}

func (per *person) GetAge() int {
	return per.age
}

func (per *person) SetSal(sal float64) { 
	if sal < 3000 || sal > 8000 {
		fmt.Println("工资范围不正确")
	}
	per.sal = sal
}

func (per *person) GetSal() float64 {
	return per.sal
}

