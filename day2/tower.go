package main

import (
	"fmt"
)

func main()  { 
	//半金字塔
	var n int = 6
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//正金字塔
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//倒金字塔
	for i := n ; i >= 1 ; i-- {
		for j := 1 ; j <= n - i ; j++ {
			fmt.Print(" ")
		} 
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//菱形
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	for i := n-1 ; i >= 1 ; i-- {
		for j := 1 ; j <= n - i ; j++ {
			fmt.Print(" ")
		} 
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//空心正金字塔
	for i := 1 ; i <= n ; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ { 
			if i== 1 || i == n {
				fmt.Print("*")
			} else {
				if k == 1 || k == (i - 1) * 2 + 1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//空心倒金字塔
	for i := n ; i >= 1 ; i-- {
		for j := 1 ; j <= n - i ; j++ {
			fmt.Print(" ")
		} 
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			if i== 1 || i == n {
				fmt.Print("*")
			} else {
				if k == 1 || k == (i - 1) * 2 + 1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
	fmt.Println("================================")

	//空心菱形
	for i := 1 ; i <= n; i++ {
		for j := 1; j <= n - i; j++ {
			fmt.Print(" ")
		}
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			if k == 1 || k == (i - 1) * 2 + 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}

	for i := n-1 ; i >= 1 ; i-- {
		for j := 1 ; j <= n - i ; j++ {
			fmt.Print(" ")
		} 
		for k := 1 ; k <= (i - 1) * 2 + 1 ; k++ {
			if k == 1 || k == (i - 1) * 2 + 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("================================")
}