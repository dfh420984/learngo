package main

import (
	"fmt"
	"learngo/day4/exec6"
)

func main()  {
	var res float64
	res = exec6.Cal(4.0, 2.0, '+')
	fmt.Println("res=", res)
}