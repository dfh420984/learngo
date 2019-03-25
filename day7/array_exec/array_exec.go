package main

import (
	"fmt"
)

func main()  {
	var hens [3]float64
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	averageWeight := fmt.Sprintf("%.2f",totalWeight / float64(len(hens)))
	fmt.Printf("鸡的总重量是%.2f,平均重量是%v", totalWeight, averageWeight)
}