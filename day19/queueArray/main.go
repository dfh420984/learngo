package main

import (
	"fmt"
	"errors"
	"os"
)

type Queue struct {
	front int
	real int 
	maxSize int
	array [5]int
}

func (this *Queue) AddQueue(val int) (err error) {
	if this.real == this.maxSize - 1 {
		return errors.New("queue full")
	}
	this.real++
	this.array[this.real] = val
	return
}

func (this *Queue) GetQueue() (val int,err error) {
	if this.real == this.front {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return
}

func (this *Queue) ShowQueue() {
	fmt.Println("当前队列情况:")
	for i := this.front+1; i <= this.real; i++ {
		fmt.Printf("array[%d]=%d\t\n",i,this.array[i])
	}
}

func main() {
	queue := &Queue{
		front : -1,
		real : -1,
		maxSize : 5,
	}
	for {
		fmt.Println("1.输入add添加到队列")
		fmt.Println("2.输入get从队列获取数据")
		fmt.Println("3.输入show显示队列")
		fmt.Println("4.输入exit退出循环")
		var key string
		var val int
		fmt.Scanln(&key)
		switch key {
			case "add":
				fmt.Println("请输入添加到队列的值")
				fmt.Scanf("%d\n",&val)
				queue.AddQueue(val)
			case "get":
				fmt.Println(queue.GetQueue())
			case "show":
				queue.ShowQueue()
			case "exit":
				os.Exit(0)
			default:
				fmt.Println("输入有误，请从新输入")
		}	
	}
}