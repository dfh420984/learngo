package main

import (
	"fmt"
	"errors"
	"os"
)

type CycleQueue struct {
	head int
	tail int
	maxSize int
	array [5]int
}

//判断队列是否已满
func (this *CycleQueue) IsFull() bool {
	if (this.tail + 1) % this.maxSize == this.head {
		return true
	} else {
		return false
	}
}

//判断队列是否为空
func (this *CycleQueue) IsEmpty() bool {
	return this.head == this.tail
}

//入队列
func (this *CycleQueue) Push(val int) (err error) {
	//1.先判断队列是否以满
	if this.IsFull() {
		return errors.New("array full")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

//出队列
func (this *CycleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return -1, errors.New("array empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

//显示队列元素个数
func (this *CycleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

//显示队列
func (this *CycleQueue) ListQueue() {
	size := this.Size()
	fmt.Println("当前队列元素个数:",size)
	if size == 0 {
		fmt.Println("empty array")
	}
	//是设计一个临时变量指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\n",tempHead,this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
}

//环形队列curd
func main()  {
	cycleQueue := &CycleQueue{
		head : 0,
		tail : 0,
		maxSize : 5,
	}
	for {
		fmt.Println("1.输入push添加到队列")
		fmt.Println("2.输入pop从队列获取数据")
		fmt.Println("3.输入list显示队列")
		fmt.Println("4.输入exit退出循环")
		var key string
		var val int
		fmt.Scanln(&key)
		switch key {
			case "push":
				fmt.Println("请输入添加到队列的值")
				fmt.Scanf("%d\n",&val)
				cycleQueue.Push(val)
			case "pop":
				fmt.Println(cycleQueue.Pop())
			case "list":
				cycleQueue.ListQueue()
			case "exit":
				os.Exit(0)
			default:
				fmt.Println("输入有误，请从新输入")
		}	
	}
}