package main

import(
	"fmt"
	"reflect"
)

//对int反射
func testRefInt(b interface{}) {
	// rType := reflect.TypeOf(b)
	// fmt.Println("rType=",rType)
	// rVal := reflect.ValueOf(b)
	// fmt.Println("rVal=",rVal)
	// n2 := 2 + rVal.Int()
	// fmt.Printf("rType=%v,rVal=%v,n2=%v \n",rType, rVal, n2)
	// iv := rVal.Interface()
	// num2 := iv.(int)
	// fmt.Println("num2=",num2)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=",rVal)
	rVal.Elem().SetInt(20)
	fmt.Println("rVal=",rVal)
}

func testRefFloat64(b interface{}) { 
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	rKind := rVal.Kind()
	rKindStr := rKind.String()
	fmt.Printf("rType=%v,rVal=%v,rKind=%v \n",rType,rVal,rKind)
	fmt.Printf("rKindStr=%v \n",rKindStr)
}

//对结构体反射
type Student struct {
	Name string
	Age int 
}

func testRefStruct(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=",rType)
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=",rVal)
	iv := rVal.Interface()
	fmt.Printf("iv=%v,iv type=%T\n",iv, iv)
	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("stu.Name=%v \n", stu.Name)
	} else {
		fmt.Printf("断言失败，stu=%v,ok=%v\n", stu, ok)
	}
}

func main()  {
	// var num int = 10
	// testRefInt(&num)
	// fmt.Println("num=",num)
	// stu := Student{Name:"jack", Age:10}
	// testRefStruct(stu)
	testRefFloat64(1.23)
}