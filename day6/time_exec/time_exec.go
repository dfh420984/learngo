package main

import (
	"fmt"
	"time"
	"strconv"
)

func test03() {
	str := ""
	for i := 0; i < 10000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	now := time.Now()
	fmt.Printf("now=%v,type=%T \n", now, now)

	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	timeStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Printf("当前时间是%v \n", timeStr)

	fmt.Printf("当前unxix时间戳%v,纳秒时间戳\n", now.Unix(), now.UnixNano())

	//1.需求每隔1s打印一个数字，打印到10退出
	//2.需求每隔0.1s打印一个数字，打印到10退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//time.Sleep(time.Second)
		time.Sleep(time.Millisecond * 100)
		if i== 10 {
			break
		}
	}


	//统计test03代码执行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("test03耗时%v秒", end-start)
}