package main

import(
	"fmt"
	"strconv"
	"math/rand"
)

//intchan
func test1() {
	intChan := make(chan int,3) 
	intChan<- 1
	intChan<- 2
	intChan<- 3
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Println(num1,num2,num3)
}

//mapchan
func test2() {
	mapChan := make(chan map[string]string, 2)
	m1 := make(map[string]string)
	m2 := make(map[string]string)
	m1["city1"] = "北京"
	m1["city2"] = "上海"
	mapChan<- m1
	m2["hero1"] = "孙悟空"
	m2["hero2"] = "貂蝉"
	mapChan<- m2
	close(mapChan)
	for {
		v, ok := <-mapChan
		if !ok {
			break
		}
		fmt.Printf("v=%v,ok=%v\n",v,ok)
	}
}

//catChan
type Cat struct {
	Name string
	Age int
}
func test3() {
	catChan := make(chan Cat, 2)
	cat1 := Cat{
		Name : "小白",
		Age : 1,
	}
	cat2 := Cat{
		Name : "小黑",
		Age : 2,
	}
	catChan<- cat1
	catChan<- cat2
	close(catChan)
	for {
		v, ok := <-catChan
		if !ok {
			break
		}
		fmt.Printf("v=%v,ok=%v\n",v,ok)
	}
}

//allChan
func test4() {
	allChan := make(chan interface{}, 10)
	allChan<- 1
	m1 := make(map[string]string, 2)
	m1["name"] = "dfh420984"
	m1["address"] = "beijing"
	allChan<- m1
	cat1 := Cat{
		Name : "小白",
		Age : 1,
	}
	allChan<- cat1
	slice := make([][2]int,2)
	slice[0] = [2]int{1,2}
	//slice = append(slice,slice1)
	slice[1] = [2]int{3,4}
	//slice = append(slice,slice2)
	allChan<- slice
	mapSlice := make([]map[string]string, 2)
	m2 := make(map[string]string,2)
	m2["name"] = "xiaoming"
	m2["address"] = "beijing"
	//mapSlice = append(mapSlice, m2)
	mapSlice[0] = m2
	m3 := make(map[string]string,2)
	m3["name"] = "tomng"
	m3["address"] = "tianjin"
	//mapSlice = append(mapSlice, m3)
	mapSlice[1] = m3
	allChan<- mapSlice
	close(allChan)
	for {
		v, ok := <-allChan
		if !ok {
			break
		}
		fmt.Printf("v=%v,ok=%v\n",v,ok)
	}
}

type Person struct {
	Name string
	Age int
	Address string
}

func test5() { 
	perChan := make(chan Person, 10)
	for i := 1; i <= 10; i++ {
		p := Person{
			Name : "stu"+strconv.Itoa(i),
			Age : rand.Intn(10),
			Address : "beijing",
		}
		perChan<- p
	}
	close(perChan)
	for {
		v, ok := <-perChan
		if !ok {
			break
		}
		fmt.Printf("v=%v,ok=%v\n",v,ok)
	}
}

func writeData(intChan chan int) {
	for i:= 1; i <= 50; i++ {
		intChan<- i
		fmt.Println("writeData", i)
	}
	close(intChan)
}

func readData(intChan chan int , exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData", v)
	}
	exitChan<- true
	close(exitChan)
}

func main()  {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		v, ok := <-exitChan 
		fmt.Printf("v=%v,ok=%v\n",v,ok)
		if !ok {
			break
		}
	}
}