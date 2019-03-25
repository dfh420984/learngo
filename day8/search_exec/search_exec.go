package main

import (
	"fmt"
)

func test() {
	var str [4]string = [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王","青衣福王"}
	var inputName string
	fmt.Println("请输入一个名称")
	fmt.Scanln(&inputName) 
	for i, val := range str {
		if inputName == val {
			fmt.Printf("找到%v,下标%v\n",val,i)
			break
		}
		if i == len(str) - 1 {
			fmt.Printf("没有找到%v\n",inputName)
		}
	}
}

//二分查找
func binary_search(arr [6]int, findVal int)  {
	low := 0
	high := len(arr) - 1
	for ((low <= high) && (low <= len(arr) - 1) && (high <= len(arr) - 1)) {
		middle := (low + high) / 2
		if findVal == arr[middle] {
			fmt.Printf("找到了下标为%v\n", middle)
			break
		} else if (findVal < arr[middle]) {
			high = middle -1
		} else {
			low = middle + 1
		}
	}
	fmt.Printf("找不到该值%v \n", findVal)
}

func main()  {
	//test()
	var arr [6]int = [6]int{100, 103, 106, 200, 201, 300} 
	binary_search(arr, 103)
}