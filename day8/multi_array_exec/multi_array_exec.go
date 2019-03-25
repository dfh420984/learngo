package main

import (
	"fmt"
)

func test() { 
	//直接赋值
	// var arr [2][3]int = [2][3]{{10,0,0},{20,0,0}}
	var arr [2][3]int
	arr[0][0] = 10
	arr[1][0] = 20
	fmt.Printf("arr[0]得地址是%p\n", &arr[0])
	fmt.Printf("arr[1]得地址是%p\n", &arr[1])
	fmt.Printf("arr[0][0]得地址是%p\n", &arr[0][0])
	fmt.Printf("arr[1][0]得地址是%p\n", &arr[1][0])
}

//for-range二维数组遍历
func test02() {
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j :=0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d]=%v \n",i, j, arr[i][j])
		}
	}

	for i, val := range arr {
		for j, val2 := range val {
			fmt.Printf("arr[%d][%d]=%v \n",i, j, val2)
		}
	}
}

//二维数组练习
func test03() {
	var score [2][3]float64
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[i]); j++ {
			fmt.Printf("请输入第%d班，第%d个学生得成绩\n", i+1, j+1)
			fmt.Scanln(&score[i][j])
		}
	}
	//统计所有班级得总分
	total := 0.0
	for _, val := range score {
		for _, val2 := range val {
			total += val2
		}
	}
	fmt.Printf("所有班级得总分%v\n", total)
}

func main()  {
	//test()
	//test02()
	test03()
}