package model

import (
	"fmt"
)
type Monkey struct {
	Name string
}

type LittleMonkey struct {
	Monkey
}

type Bird interface {
	Fly()
}

type Fish interface {
	Swim()
}

func (this *Monkey) Climb() {
	fmt.Println(this.Name, "会爬树")
}

func (this *LittleMonkey) Fly() {
	fmt.Println(this.Name, "通过学习会飞行")
}

func (this *LittleMonkey) Swim() {
	fmt.Println(this.Name, "通过学习会游泳")
}
