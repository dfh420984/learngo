package main

import (
	"fmt"
	"math/rand"
	"time"
)
func test() {
	var hens [3]float64
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	averageWeight := fmt.Sprintf("%.2f",totalWeight / float64(len(hens)))
	fmt.Printf("鸡的总重量是%.2f,平均重量是%v\n", totalWeight, averageWeight)
}

func test01() {
	var score [3]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个元素得值\n", i)
		fmt.Scanln(&score[i])
	}

	//数组变量打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}
}

//输出26个字符
func test02() {
	var mychar [26]byte
	for i:=0; i < 26; i++ {
		mychar[i] = 'A' + byte(i)
	}

	for i:=0; i < 26; i++ {
		fmt.Printf("%c ", mychar[i])
	}
}

//找出数组中最大的值
func test03(arr [4]int) (index int, val int) {
	var maxIndex int = 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[maxIndex] {
			maxIndex = i
		}
	}
	index = maxIndex
	val = arr[maxIndex]
	return index, val
}

//4.求一个数组的平均值
func test04(arr [4]int) float64 {
	var total int = 0
	for i:=0; i < len(arr); i++ {
		total += arr[i]
	}
	return float64(total / len(arr))
}

/**
	5.随机生成五个数,并将其反转打印
	最后一个元素和第一个元素交换，倒数第二个元素和第二个元素交换，
**/
func test05() {
	var arr [5]int
	len := len(arr)
	rand.Seed(time.Now().Unix())
	for i := 0; i< len; i++ {
		arr[i] = rand.Intn(100)
	}
	fmt.Println("交换前arr=", arr)
	temp := 0
	for i := 0; i < len / 2; i++ {
		temp = arr[i]
		arr[i] = arr[len-i-1]
		arr[len-i-1] = temp
	}
	fmt.Println("交换后arr=", arr)
}

func main()  {
	
	test()

	//test01()

	//四种数组初始化方式,数组类型由值决定
	var numArr1 [3]int = [3]int{1, 2, 3} 
	fmt.Println("numArr1=", numArr1)

	var numArr2 = [3]int{4, 5, 6} 
	fmt.Println("numArr2=", numArr2)

	var numArr3 = [...]int{7, 8, 9} 
	fmt.Println("numArr3=", numArr3)

	var numArr4 = [...]int{1:700, 2:900, 3:100} 
	fmt.Println("numArr4=", numArr4)

	var numArr5 = [...]string{1:"tom", 2:"jack", 3:"mary"} 
	fmt.Println("numArr5=", numArr5)

	test02()
	fmt.Println("")

	var numArr6 = [4]int{5,8,3,2}
	maxIndex, maxVal := test03(numArr6)
	fmt.Printf("最大值是%v,下标是%v\n", maxVal, maxIndex)

	averageVal := test04(numArr6)
	fmt.Printf("数组平均值是%v\n", averageVal)

	test05()
}