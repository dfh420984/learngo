package model

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

type Puiple struct {
	Student
}

type Graduate struct {
	Student
}

func (stu *Student) ShowInfo() {
	fmt.Printf("Name=%v, Age=%v, Score=%v \n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

func (pu *Puiple) Testing() {
	fmt.Println("小学生考试中")
}

func (gu *Graduate) Testing() {
	fmt.Println("大学生考试中")
}

