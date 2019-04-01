package main

import(
	"fmt"
	"time"
	"strconv"
	"runtime"
	"sync"
)

func test()  {
	for i := 1; i<=10; i++ {
		fmt.Println("test() hello,world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func createGoroutine() {
	go test()  //开启一个协程
	for i := 1; i<=10; i++ {
		fmt.Println("main() hello,world"+strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func getCpuNum() {
	num := runtime.NumCPU()
	fmt.Println("cpu num = ",num)
}

var (
	myMap = make(map[int]int)
	//声明一个全局的同步互斥锁
	lock sync.Mutex
)

func jisuan(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() { 
	//createGoroutine()
	//getCpuNum()
	for i:=1; i <= 10; i++ {
		go jisuan(i)
	}
	time.Sleep(time.Second*6)
	for i,v := range myMap {
		fmt.Printf("map[%d]=%d\n",i,v)
	}
}