package main

import (
	"fmt"
)

func cal(n1 float64, n2 float64, oper byte) float64 {
	var res float64
	switch oper {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2 
		case '*':
			res = n1 * n2 
		case '/':
			res = n1 / n2
		default:
			fmt.Println("操作符不合法。。")
	}
	return res	
}

func main()  {
	var res float64
	res = cal(4.0, 2.0, '+')
	fmt.Println("res=", res)
}