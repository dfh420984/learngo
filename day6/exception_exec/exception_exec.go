package main

import (
	"fmt"
	"time"
	"errors"
)

//接收异常并输出
func test() { 
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println("error=", error)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

//自定义异常并输出
func readconf(name string) error {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误。。")
	}
}

func test02() {
	err := readconf("config2.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test02继续执行")
}

func main()  {
	test()
	time.Sleep(time.Second)
	fmt.Println("test异常后面得代码继续执行")
	test02()
}