package main

import (
	"fmt"
)

//选择排序，从大到小排序
func selectSort(arr *[6]int) {
	for i := 0; i < len(arr)-1; i++ {
		//假设第一个元素是最大值
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		if maxIndex != i { //开始交换
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
}

func main() {
	arr := [6]int{10, 88, 23, 40, 55, 90}
	selectSort(&arr)
	fmt.Println(arr)
}
