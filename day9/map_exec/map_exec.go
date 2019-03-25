package main

import (
	"fmt"
	"sort"
)

//1.定义map
func test() {

	//第一种方式
	var cities map[string]string
	cities = make(map[string]string)
	cities["01"] = "beijing"
	cities["02"] = "shanghai"
	cities["03"] = "guangzhou"
	cities["04"] = "shenzheng"
	fmt.Println(cities)

	//第二种方式
	hero := make(map[string]string)
	hero["01"] = "武大郎"
	hero["02"] = "武松"
	hero["03"] = "潘金莲"
	fmt.Println(hero)

	//第三种方式
	a  := map[string]string{
		"01":"wa",
		"02":"ha",
		"03":"ha",
	}
	fmt.Println(a)
}

//2.存放三个学生信息
func test02() {
	stuMap := make(map[string]map[string]string, 3)
	stuMap["01"] = make(map[string]string, 2)
	stuMap["01"]["name"] = "xiaoming"
	stuMap["01"]["sex"] = "女"

	stuMap["02"] = make(map[string]string, 2)
	stuMap["02"]["name"] = "jack"
	stuMap["02"]["sex"] = "男"

	stuMap["03"] = make(map[string]string, 2)
	stuMap["03"]["name"] = "marry"
	stuMap["03"]["sex"] = "女"
	fmt.Println(stuMap)
}

//3.map的curd
func test03() {
	hero := make(map[string]string)
	hero["01"] = "武大郎"
	hero["02"] = "武松"
	hero["03"] = "潘金莲"
	//删除指定key
	delete(hero, "03")
	//新增
	hero["04"] = "西门庆"
	//全部删除
	//hero = make(map[string]string)
	fmt.Println(hero)
	//map查找
	v, ok := hero["02"]
	if ok {
		fmt.Println(v)
	}

}

//4.遍历map
func test04() {
	stuMap := make(map[string]map[string]string, 3)
	stuMap["stu1"] = make(map[string]string, 2)
	stuMap["stu1"]["name"] = "xiaoming"
	stuMap["stu1"]["sex"] = "女"

	stuMap["stu2"] = make(map[string]string, 2)
	stuMap["stu2"]["name"] = "jack"
	stuMap["stu2"]["sex"] = "男"

	stuMap["stu3"] = make(map[string]string, 2)
	stuMap["stu3"]["name"] = "marry"
	stuMap["stu3"]["sex"] = "女"

	for k1,v1 := range stuMap {
		fmt.Printf("k1=%v,v1=%v \n",k1,v1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v,v2=%v \n",k2,v2)
		}
	}
}

//map切片
func test05() {
	var monster []map[string]string
	monster = make([]map[string]string, 2)
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["name"] = "孙悟空"
		monster[0]["age"] = "6000"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string, 2)
		monster[1]["name"] = "组八戒"
		monster[1]["age"] = "1000"
	}

	//动态增加monster
	newMonster := map[string]string {
		"name":"沙和尚",
		"age":"300",
	}
	monster = append(monster, newMonster)
	fmt.Println(monster)
}

//map排序
func test06() {
	map1 := make(map[int]int, 10)
	map1[6] = 60
	map1[8] = 10
	map1[3] = 80
	map1[2] = 20

	var keys []int

	for k, _ := range map1 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(map1[v])
	}
}

//map课堂练习
func modifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["pwd"] = "888888"
	} else {
		user[name] = make(map[string]string, 2)
		user[name]["nick"] = "nick" + name
		user[name]["pwd"] = "888888"
	}
}

func main()  {
	//test()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	user := make(map[string]map[string]string)
	user["mali"] = make(map[string]string)
	user["mali"]["nick"] = "nickmali"
	user["mali"]["pwd"] = "123"
	modifyUser(user, "xiaoming")
	modifyUser(user, "mali")

	fmt.Println(user)
}