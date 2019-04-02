package main

import(
	"fmt"
	_"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan<- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool 
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num % i == 0 { //不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
		}
	}
	exitChan<- true
}

func main()  { 
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go putNum(intChan)
	//开启4个协程处理数据
	for i := 1; i<=8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	go func(){
		for i := 1; i<=8; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()
	for {
		res,ok := <-primeChan
		if !ok {
			break
		}
		fmt.Printf("素数=%d \n", res)
	}
	fmt.Println("main主线程退出")
}