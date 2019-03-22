package main

import (
	"fmt"
)

func main()  {
	//lable2:
	for i := 1; i <= 4; i++ {
		//lable1:
		for j :=1; j <= 4; j++ {
			if j == 2 {
				//break
				//break lable1
				//break lable2
				//continue lable1
				//continue lable2
				continue
			}
			fmt.Println("j=", j)
		}
	}
}