package main

import "fmt"

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key { //在基准点左边找出比基准点大的数
				i++
			}
			for arr[j] > key { //在基准点右边找出比基准点小的数
				j--
			}
			if i <= j { //找到后交换
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		//数组左右根据基准点交换完毕后，左右半边开始递归
		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

func main() {
	arr := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
