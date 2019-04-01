package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"io"
	"io/ioutil"
	"flag"
	"encoding/json"
)

func test01() {
	file, err := os.Open("C:/Users/duanfuhao/go/src/learngo/day14/test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("打开文件成功")
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}

//buffio缓冲写文件
func test02() {
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day14/test.txt", os.O_RDWR | os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	str := "hello world \r\n"
	//带缓冲写入
	writer := bufio.NewWriter(file)
	for i :=0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

//buffio读文件字节
func test03() {
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day14/test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	p := make([]byte, 100)
	n , err := reader.Read(p)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("n=",n)
}

//buffio读取文件字符串
func test04() {
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day14/test.txt", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for { 
		line, err := reader. ReadString('\n')
		if err ==  io.EOF {
			break
		}
		fmt.Print(line)
	}
}

func test05() {
	data, err := ioutil.ReadFile("C:/Users/duanfuhao/go/src/learngo/day14/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("C:/Users/duanfuhao/go/src/learngo/day14/test2.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}


//检查文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//文件拷贝
func CopyFile(dstFileName string, srcFileName string) (written int64, err error){
	srcFil, err := os.Open(srcFileName)
	defer srcFil.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	reader := bufio.NewReader(srcFil)
	dstFile, err :=  os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	return io.Copy(writer, reader)
}

func test06() {
	srcFileName := "C:/Users/duanfuhao/go/src/learngo/day14/test.txt"
	dstFileName := "C:/Users/duanfuhao/go/src/learngo/day14/test3.txt"
	_, err := CopyFile(dstFileName, srcFileName)
	if err != nil {
		fmt.Printf("拷贝错误%v",err)
	} else {
		fmt.Println("拷贝完成")
	}
}

//统计字符个数案例
func test07() {
	type CountStr struct {
		EnCount int //记录英文个数
		NumCount int //记录数字个数
		SpaceCount int //记录空格个数
		OtherCount int //记录其他字符个数
	}
	count := CountStr{}
	file, err := os.OpenFile("C:/Users/duanfuhao/go/src/learngo/day14/test.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("文件打开错误%v",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str1 := []rune(str)
		for _, v := range str1 {
			switch  {
				case v >='a' && v<='z':
					fallthrough
				case v >='A' && v<='Z':
					count.EnCount++
				case v == '\t' || v == '\n':
					count.SpaceCount++
				case v >='0' && v<='9':
					count.NumCount++
				default:
					count.OtherCount++
				
			}
		}
	}
	fmt.Println(count)

}

//获取命令行参数
func test08() {
	for i,v := range os.Args {
		fmt.Printf("index=%v,val=%v \n",i,v)
	}
}

//flag包解析命令行参数 如 main.exe -u root -pwd root -h 192.168.1.1 -port 3306
func test09() {
	var user string
	var pwd string
	var host string 
	var port int
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "主机默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口默认为3306")
	flag.Parse()
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, pwd, host, port)
}

//结构体，map,slice序列化json
type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

//结构体，map,切片序列化
func test10()  { 
	//结构体
	monster := &Monster{
		Name : "孙悟空",
		Age : 2000,
		Birthday : "1980-06-13",
		Sal : 20000,
		Skill : "筋斗云",
	}

	//map
	mapdata := make(map[string]interface{})
	mapdata["name"] = "红孩儿"
	mapdata["age"] = 200
	mapdata["address"] = "火云洞"

	//切片
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "猪八戒"
	m1["age"] = 5000
	m1["address"] = [2]string{"洛水河","高老庄"}
	slice = append(slice, m1)
	m2 := make(map[string]interface{})
	m2["name"] = "唐僧"
	m2["age"] = 6000
	m2["address"] = [2]string{"长安","珞珈山"}
	data, _ := json.Marshal(monster)
	data2, _ := json.Marshal(mapdata)
	data3, _ := json.Marshal(slice)
	fmt.Printf("monster序列化后=%v\n", string(data))
	fmt.Printf("mapdata序列化后==%v\n", string(data2))
	fmt.Printf("slice序列化后=%v\n", string(data3))
}

//json反序列化
func test11() {
	monster := &Monster{
		Name : "孙悟空",
		Age : 2000,
		Birthday : "1980-06-13",
		Sal : 20000,
		Skill : "筋斗云",
	}
	data, _ := json.Marshal(monster)
	json.Unmarshal(data, monster)
	fmt.Printf("反序列化后=%v\n", monster)
}

func main()  {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	//test09()
	//test10()
	test11()
}