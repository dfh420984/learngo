package main

import (
	"fmt"
)

//冒泡排序,从小到大排序
func bubleSort(arr *[6]int) { 
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1],arr[j] = arr[j],arr[j+1]
			}
		}
	}
	fmt.Println(arr)
}

func main()  {
	arr := [6]int{10,88,23,40,55,90}
	bubleSort(&arr)
	fmt.Println(arr)
}