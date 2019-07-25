package main

import (
	"fmt"
	"regexp"
)

const myEmail = `
this is my email 352928736@qq.com
email1 abc@123.com
email2 haeea@org.com
email3 duanfuhao@smzdm.com.cn
`

func main() {
	reg := regexp.MustCompile(`(\w+)@(\w+)(\.[\w.]+)`)
	//str := reg.FindString(myEmail)
	//str := reg.FindAllString(myEmail, -1)
	str := reg.FindAllStringSubmatch(myEmail, -1)
	fmt.Printf("%s\n", str)
	for _, v1 := range str {
		for _, v2 := range v1 {
			fmt.Printf("%s\n", v2)
		}
		fmt.Println()
	}
}
