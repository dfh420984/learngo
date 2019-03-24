package main

import (
	"fmt"
)

//fib 斐波拉切
//1 1 2 3 5 8
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

//猴子吃桃
//第十天桃子 = 1，第9天桃子 = （第十天桃子 + 1） * 2
func peach(n int) int {
	if n == 10 {
		return 1
	} else {
		return (peach(n + 1) + 1) * 2
	}
}

//毕包函数
func wrapper(n1 int) func (int) (int) {
	res1 := n1 
	fmt.Println("res1=",res1)
	return func(x int) int {
		res2 := n1 + x
		return res2
	}
}



func main()  {
	// res1 := fib(1)
	// res2 := fib(2)
	// res3 := fib(3)
	// fmt.Printf("res1=%v,res2=%v,res=%v \n",res1,res2,res3)

	//第一天桃子数量是
	res1 := peach(1)
	fmt.Printf("第一天桃子数量是%v \n", res1)

	//匿名函数
	fun1 := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Printf("n1 + n2 = %v \n", fun1(1, 2))

	//比保钓用
	bibao := wrapper(10)
	bbres := bibao(10)
	fmt.Println("必报结果是", bbres)

}