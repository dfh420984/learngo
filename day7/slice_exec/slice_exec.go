package main

import (
	"fmt"
)

//1.通过数组取出切片
func test() {
	var arr [5]int = [5]int{11, 22, 33, 44, 55}
	slice := arr[1:3]
	fmt.Println("slice得元素是=",slice)
	fmt.Println("slice得元素个数是=",len(slice))
	fmt.Println("slice得容量是=",cap(slice))
}

//2.通过make创建切片
func test01() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10.0
	slice[3] = 30.0
	slice[4] = 60.0
	fmt.Println("test01 slice得元素是=",slice)
	fmt.Println("test01 slice得元素个数是=",len(slice))
	fmt.Println("test01 slice得容量是=",cap(slice))
}

//3.定义一个切片，直接指定具体数组
func test03() {
	var slice []string = []string{"hello", "world", "ok"}
	fmt.Println("test03 slice得元素是=",slice)
	fmt.Println("test03 slice得元素个数是=",len(slice))
	fmt.Println("test03 slice得容量是=",cap(slice))
}

//4.切片遍历,切片元素得值改变后，所有在相关数组或切片相同得值都得改变
func test04() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%d]=%v\n",i,slice[i])
	// } 
	
	for i,vs := range slice {
		fmt.Printf("slice[%d]=%v\n",i,vs)
	}

	slice2 := slice[1:2]
	slice2[0] = 100
	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)
}

//5.append对切片动态增加
func test05() {
	var slice []int = []int {100, 200, 300}
	slice = append(slice, 400, 500)
	fmt.Println("slice=",slice )
	slice = append(slice, slice...)
	fmt.Println("slice=",slice )
}

//6.切片copy操作
func test06() {
	var slice1 []int = []int{1,2,3,4,5}
	var slice2 []int = make([]int, 10)
	copy(slice2, slice1)
	//slice2[0] = 100
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2=", slice2)
}

//7.string底层是[]byte，可以进行切片操作
func test07() {
	var str string = "hello world"
	slice := str[1:]
	fmt.Println("slice=", slice)

	//修改字符串
	arr1 := []byte(str)
	arr1[0] = 'H'
	str = string(arr1)
	fmt.Println("str=", str)

	arr2 := []rune(str)
	arr2[0] = '好'
	str = string(arr2)
	fmt.Println("str=", str)
}

//8.切片课堂练习
func fbn(n int) ([]int64) { 
	fbnslice := make([]int64, n)
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < n; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}

func main()  {
	test()
	test01()
	test03()
	test04()
	test05()
	test06()
	test07()
	fnbslice := fbn(10)
	fmt.Println("fnbslice=", fnbslice)
}