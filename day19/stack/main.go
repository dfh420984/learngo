package main

import (
	"fmt"
	"errors"
)

type Stack struct {
	maxTop int //表示栈最大可以存放个数
	Top int //表示栈顶
	arr [5]int  //数组模拟栈
}

//入栈
func (this *Stack) push(val int) (err error) {
	//先判断栈是否满了
	if this.Top == this.maxTop - 1 { 
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

//出栈
func (this *Stack) pop() (val int, err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return -1,errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

//显示栈
func (this *Stack) list() (err error) {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return errors.New("stack empty")
	}
	for i:= this.Top; i>=0; i-- {
		fmt.Printf("arr[%d]=%d\n",i,this.arr[i])
	}
	return
}

func main()  {
	stack := &Stack{
		maxTop : 5,
		Top : -1,
	}
	stack.push(10)
	stack.push(20)
	stack.list()
}