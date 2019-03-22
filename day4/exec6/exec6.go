package exec6

import (
	"fmt"
)

func Cal(n1 float64, n2 float64, oper byte) float64 {
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