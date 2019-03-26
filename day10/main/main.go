package main

import (
	"fmt"
	"learngo/day10/model"
)

func main()  {
	stu := model.Student{
		Name : "tom",
		Score : 96.0,
	}
	fmt.Println(stu)

	stu2 := model.Newstudent("jack", 100.0)
	fmt.Println(*stu2)
	fmt.Println((*stu2).GetScore())
}