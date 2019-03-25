package main

import (
	"fmt"
)

//冒泡排序
func bubble_sort(arr [6]int) ([6]int) {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j + 1]
				arr[j + 1] = tmp
			}
		}
	}
	return arr
}

func bubble_sort_ptr(arr *[6]int) {
	for i := 0; i < len(*arr) - 1; i++ {
		for j := 0; j < len(*arr) - i - 1; j++ {
			if arr[j] > (*arr)[j+1] {
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j + 1] = tmp
			}
		}
	}
}

func main()  {
	arr := [6]int{2,6,4,8,7,1}
	//arr = bubble_sort(arr)
	bubble_sort_ptr(&arr)
	fmt.Println("arr=",arr)
}