package main

import (
	"fmt"
	"reflect"
)

type ReflectDemo struct {
	Age  int    `json:"age" gorm:"age"`
	Name string `json:"name" gorm:"name"`
	Addr string `json:"addr" gorm:"addr"`
}

func (r *ReflectDemo) SetAge(age int) {
	r.Age = age
}

func (r *ReflectDemo) SetName(name string) {
	r.Name = name
}

func (r *ReflectDemo) SetAddr(addr string) {
	r.Addr = addr
}

// 反射相关操作
func main() {
	//获取relect.Value, relect.Type
	//反射结构体指针和结构体是不一样,如果是指针，需要调用Elem()
	v := reflect.ValueOf(&ReflectDemo{})
	vElem := v.Elem()
	t := v.Type()
	tElem := t.Elem()
	//获取结构体中方法
	numMethod := v.NumMethod()
	fmt.Println("numMethod:", numMethod)
	for i := 0; i < numMethod; i++ { //动态调用方法并给属性复制
		vobj := v.Method(i)
		tobj := t.Method(i)
		fmt.Printf("method name : %s, \n", tobj.Name)
		if tobj.Name == "SetAge" {
			vobj.Call([]reflect.Value{reflect.ValueOf(100)})
		}
		if tobj.Name == "SetName" {
			vobj.Call([]reflect.Value{reflect.ValueOf("zhangshan")})
		}
		if tobj.Name == "SetAddr" {
			vobj.Call([]reflect.Value{reflect.ValueOf("北京")})
		}
	}
	//获取结构体字段
	numField := tElem.NumField()
	fmt.Println("numField:", numField)
	for i := 0; i < numField; i++ {
		fmt.Printf("type:%s, field:%s, value: %v \n", tElem.Field(i).Type.Kind(), tElem.Field(i).Name, vElem.FieldByIndex([]int{i}))
		tagStrJson := tElem.Field(i).Tag.Get("json")
		tagStrGorm := tElem.Field(i).Tag.Get("json")
		fmt.Println("tagStr json:", tagStrJson)
		fmt.Println("tagStr gorm:", tagStrGorm)
	}
}
