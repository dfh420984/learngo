package main

import(
	"fmt"
	"time"
)

func sayHello() {
	for i:= 0; i < 10; i++ {
		fmt.Println("sayHello", i)
	}
}

func test() {
	defer func(){
		if err := recover(); err !=nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	var m1 map[string]string
	m1["name"] = "jack"
}

func main(){
	go sayHello()
	go test()
	for i := 0; i < 3; i++ {
		fmt.Println("main() ok", i)
		time.Sleep(time.Second)
	}
}