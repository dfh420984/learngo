package main

import "fmt"

func main() {
	// var char byte
	// fmt.Println("请输入一个字符")
	// fmt.Scanf("%c",&char)
	// switch char {
	// 	case 'a' :
	// 		fmt.Println("A")
	// 	case 'b' :
	// 		fmt.Println("B")
	// 	case 'c' :
	// 		fmt.Println("C")
	// 	case 'd' :
	// 		fmt.Println("D")
	// 	default :
	// 		fmt.Println("other")
	// }

	// var score float64
	// fmt.Println("请输入一个成绩")
	// fmt.Scanln(&score)
	// switch  {
	// 	case score >= 60 && score < 80 :
	// 		fmt.Println("及格")
	// 	case score >= 80 && score <= 100 :
	// 		fmt.Printf("优秀,成绩是%v \n",score)
	// 	case score < 60 :
	// 		fmt.Println("不及格")
	// 	default :
	// 		fmt.Println("不合法输入")
		
	// } 

	var month byte
	fmt.Println("请输入月份")
	fmt.Scanln(&month)
	switch month {
		case 3,4,5 :
			fmt.Println("春季")
		case 6,7,8 :
			fmt.Println("夏季") 
		case 9,10,11 :
			fmt.Println("秋季")
		case 12,1,2 :
			fmt.Println("冬季")
		default :
			fmt.Println("输入错误")
	}
}