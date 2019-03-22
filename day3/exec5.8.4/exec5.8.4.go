package main

import (
	"fmt"
)

func main()  {
	var money float64 = 100000
	var count int = 0
	for {
		if money > 50000 {
			money = money - money * 0.05
		} else if (money >= 1000 && money <= 50000) {
			money = money - 1000
		} else {
			break
		}
		count++
	}
	fmt.Printf("该人可以经过路口%d次", count)
}