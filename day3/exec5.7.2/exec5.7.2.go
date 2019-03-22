package main

import (
	"fmt"
)

func main()  { 
	var num1 int = 2
	var num2 int = 2
	var allTotal float64 = 0.0
	var passcount int = 0
	for i := 1 ; i <= num1 ; i++ { 
		var everyScore float64 = 0.0
		for j := 1; j <= num2; j++ { 
			var score float64 = 0.0
			fmt.Printf("请输入第%d个班，第%d个学生得成绩 \n", i, j)
			fmt.Scanln(&score)
			everyScore += score 
			if score >= 60 {
				passcount++
			}
		}
		fmt.Printf("请输入第%d个班得平均成绩是%f \n", i, everyScore / float64(num2))
		allTotal += everyScore
	}
	fmt.Printf("总成绩是%v,所有班得平均成绩是%v,总及格人数是%d \n", allTotal, allTotal / float64(num2), passcount)
}