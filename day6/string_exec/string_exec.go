package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main()  {
	//1.统计字符串长度
	str := "hello"
	fmt.Println("str len=", len(str))

	//2.字符串遍历，处理有中文字符
	str = "hello北京"
	cnstr := []rune(str)
	for i:=0; i < len(cnstr); i++ {
		fmt.Printf("%c\n", cnstr[i])
	}

	//3.字符串转整数
	n,err := strconv.Atoi("123")
	fmt.Printf("n=%v,err=%v,type=%T \n",n,err,n)

	//4.整数转字符串
	str = strconv.Itoa(123)
	fmt.Printf("str=%v,type=%T \n",str,str)

	//5.字符串转[]byte
	var bytes = []byte("hello world")
	fmt.Printf("bytes=%v \n", bytes)

	//6.[]byte转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v \n", str)

	//7.10进制转2，8，16进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("对应2进制%v\n", str)
	str = strconv.FormatInt(123, 8)
	fmt.Printf("对应8进制%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("对应16进制%v\n", str)

	//8.查找字符串中是否存在自定字符
	b := strings.Contains("seafool", "fool")
	fmt.Printf("b=%v\n", b)

	//9.统计一个字符串有几个指定字串
	num := strings.Count("chessss", "s")
	fmt.Printf("num=%v\n", num)

	//10.不区分大小写比较（==是区分大小写）
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b)
	fmt.Printf("b=%v\n", "abc" == "ABC")

	//11.返回字串在字符串第一次出现得位置
	index := strings.Index("wahaha", "ha")
	fmt.Printf("index=%v\n", index)

	//12.返回字串在字符串最后一次出现得位置
	index = strings.LastIndex("wahaha", "ha")
	fmt.Printf("index=%v\n", index)

	//13.字符串替换
	str = strings.Replace("go golang", "go", "go语言", 1)
	fmt.Printf("str=%v\n", str)

	//14.字符串按照指定符号拆分
	arr := strings.Split("hello,world,ok",",")
	for index, val := range arr {
		fmt.Printf("index=%v,val=%v\n", index, val)
	}

	//15.字符串大小写转换
	str = strings.ToUpper("test")
	fmt.Printf("str=%v\n", str)

	str = strings.ToLower("TEST")
	fmt.Printf("str=%v\n", str)

	//16.去掉左右两端空格
	str = strings.TrimSpace(" wahaha ")
	fmt.Printf("str=%v\n", str)

	//17.去掉左右字符
	str = strings.Trim("!wahaha!", "!")
	fmt.Printf("str=%v\n", str)

	//18.去掉左字符
	str = strings.TrimLeft("!wahaha!", "!")
	fmt.Printf("str=%v\n", str)

	//19.去掉右字符
	str = strings.TrimRight("!wahaha!", "!")
	fmt.Printf("str=%v\n", str)

	//20.判断字符串是否以自定字符串开头
	b = strings.HasPrefix("ftp://192.168.1.1", "ftp")
	fmt.Printf("b=%v\n", b)

	//21.判断字符串是否以自定字符串结尾
	b = strings.HasSuffix("ftp://192.168.1.1", "ftp")
	fmt.Printf("b=%v\n", b)

}