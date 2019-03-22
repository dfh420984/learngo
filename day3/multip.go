package main

import (
	"fmt"
)

func main()  {
	//99乘法口诀
	var n int = 9
	for i := 1 ; i <= n ; i++ {
		for j := 1 ; j <= i ; j++ {
			fmt.Printf("%d*%d=%v" , j , i , j * i) 
			fmt.Print(" ")
		} 
		fmt.Println("")
	}
}