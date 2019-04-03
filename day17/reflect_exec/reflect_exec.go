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

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float64 `json:"monster_score"`
	Sex string 
}

func (this *Monster) Jisuan(n1 int, n2 int) int {
	return n1 + n2
}

func (this *Monster) Set(name string, age int, score float64, sex string)  {
	this.Name = name
	this.Age = age
	this.Score = score
	this.Sex = sex
}

func testRefMonster(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kind := val.Kind()
	fmt.Printf("typ=%v,val=%v,kind=%v \n", typ, val, kind)
	if kind != reflect.Ptr {
		fmt.Println("expect Ptr")
		return 
	}
	//获取字段个数
	num := val.Elem().NumField()
	fmt.Println("num=", num)

	//获取字段名称和tag标签名
	for i := 0; i < num; i++ {  
		//structField := typ.Field(i)
		//获取字段名称
		fmt.Printf("typ.Field(%d).name=%v \n", i, typ.Elem().Field(i).Name)
		//获取tag标签
		fmt.Printf("typ.Field(%d).tag_name=%v \n", i, typ.Elem().Field(i).Tag.Get("json"))
	}

	//获取方法个数
	numMethod := val.NumMethod()
	fmt.Println("numMethod=", numMethod)

	//调用两个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Printf("res=%v \n", res[0].Int())
	// for i := 0; i < numMethod; i++ {  
	// 	fmt.Printf("val.Method(%d)=%v \n", i, val.Method(i))
	// }
	var params1 []reflect.Value
	params1 = append(params1, reflect.ValueOf("dfh420984"))
	params1 = append(params1, reflect.ValueOf(20))
	params1 = append(params1, reflect.ValueOf(100.0))
	params1 = append(params1, reflect.ValueOf("男"))
	val.Method(1).Call(params1)

}

func main()  {
	// var num int = 10
	// testRefInt(&num)
	// fmt.Println("num=",num)
	// stu := Student{Name:"jack", Age:10}
	// testRefStruct(stu)
	//testRefFloat64(1.23)
	monster := Monster{
		Name : "暴君",
		Age : 30,
		Score : 99.9,
		Sex : "男",
	}
	testRefMonster(&monster)
	fmt.Println(monster)
}