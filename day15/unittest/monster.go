package unittest

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int 
	Skill string
}

func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("Marshal err", err)
		return false
	}

	filePath := "C:/Users/duanfuhao/go/src/learngo/day15/test.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err", err)
		return false
	}
	return true
}

func (this *Monster) ReStore() bool {
	filePath := "C:/Users/duanfuhao/go/src/learngo/day15/test.txt"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("Unmarshal file err", err)
		return false
	}
	return true
}