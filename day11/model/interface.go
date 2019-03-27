package model

import (
	"fmt"
)

//定义usb接口
type Usb interface {
	Start()
	Stop()
}

//定义一个Phone结构体
type Phone struct {
	Name string
}

//让Phone实现Start() Stop()
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p Phone) Call() {
	fmt.Println("手机正在打电话")
}

//定义一个Cammera结构体
type Cammera struct {
	Name string
}

//让Cammera实现Start() Stop()
func (c Cammera) Start() {
	fmt.Println("相机开始工作")
}

func (c Cammera) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {

}

//usb是实现了Usb接口方法得结构体变量
func (com Computer) Working(usb Usb) {
	usb.Start()
	//接口类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

type A1 interface {
	Say()
}

type B1 interface {
	Hello()
}

type D1 interface { 
	A1
	B1
	Hi()
}

type C1 struct {
	
}

func (c C1) Say() {
	fmt.Println("say")
}

func (c C1) Hello() {
	fmt.Println("Hello")
}

func (c C1) Hi() {
	fmt.Println("Hi")
}

type E1 interface {
	Say()
}

type F1 struct {

}

func (f *F1) Say() {
	fmt.Println("f1 say")
}