package main

import (
	"fmt"	
)

//插入排序
func insertSort (arr *[6]int) { 
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i] //保存要插入的值
		insertIndex := i - 1 //插入的下标，
		//从小到大插入
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex + 1] = arr[insertIndex] //数据后移，复制一份
			insertIndex--  //指标前移
		}
		//交换数据
		if insertIndex + 1 != i {
			arr[insertIndex + 1] = insertVal
		}
	}
}

func main()  {
	arr := [6]int{10,88,23,40,55,90}
	insertSort(&arr)
	fmt.Println(arr)
}