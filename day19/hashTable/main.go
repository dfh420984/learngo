package main

import (
	"fmt"
	"os"
)

//雇员结构体
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

//雇员链表结构体
type EmpLink struct {
	Head *Emp
}

//hashtable结构体
type HashTable struct {
	LinkArr [7]EmpLink
}

//插入雇员
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   //创建一个指向当前头节点的辅助指针
	var pre *Emp = nil //创建一个pre指针在cur之前
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是一个空链表需找到对应地方插入(编号从小到大)
	for {
		if cur != nil {
			if cur.Id > emp.Id {
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}

//显示链表信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("%d是一个空链表\n", no)
		return
	}
	cur := this.Head
	for {
		if cur == nil {
			break
		}
		fmt.Printf("链表%d的雇员信息:(雇员id=%d,名字=%s) \n", no, cur.Id, cur.Name)
		cur = cur.Next
	}
}

//给hashTable编写添加雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数确认雇员添加到哪个链表
	linkNo := this.HashFunc(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}

//显示所有链表雇员信息
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//查找雇员
func (this *HashTable) Search(no int) {
	flag := false
JLoop:
	for i, v := range this.LinkArr {
		cur := v.Head
		for {
			if cur == nil {
				break
			}
			if cur.Id == no {
				fmt.Printf("链表%d的雇员信息:(雇员id=%d,名字=%s) \n", i, cur.Id, cur.Name)
				flag = true
				break JLoop
			}
			cur = cur.Next
		}
	}
	if !flag {
		fmt.Println("该雇员不存在")
		return
	}
}

func (this *HashTable) HashFunc(id int) int {
	return id % 7
}

func main() {
	var key int
	id := 0
	name := ""
	var hashTable HashTable
	for {
		fmt.Println("雇员关系管理系统")
		fmt.Println("1.添加雇员")
		fmt.Println("2.显示雇员")
		fmt.Println("3.查找雇员")
		fmt.Println("4.退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("请输入雇员id")
			fmt.Scanf("%d\n", &id)
			fmt.Println("请输入雇员姓名")
			fmt.Scanf("%s\n", &name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case 2:
			hashTable.ShowAll()
		case 3:
			fmt.Println("请输入雇员id")
			fmt.Scanf("%d\n", &id)
			hashTable.Search(id)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
