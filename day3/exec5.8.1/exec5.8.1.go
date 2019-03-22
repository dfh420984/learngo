package main

import (
	"fmt"
)

func main() {
	var n int = 100
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("当和第一次大于20得数是", i)
			break
		}
	} 
}