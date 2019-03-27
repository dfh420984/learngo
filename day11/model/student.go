package model

import (
	"fmt"
)

type Student struct {
	Name string
	Age int
	Score int
}

//声明一个Student类型切片，并按Score从小到大排序
type StuSlice []Student

//实现sort.Sort(data Interface)接口方法
func (hs StuSlice) Len() int {
	return len(hs)
}

func (hs StuSlice) Less(i, j int) bool {
	return hs[i].Score < hs[j].Score
}

func (hs StuSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
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

