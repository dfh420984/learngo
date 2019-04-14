package main

import (
	"fmt"
	"errors"
)

//单向链表
type HeroNode struct {
	no int
	name string
	nickname string
	next *HeroNode
}

//插入链表(在链表最后插入)
func (this *HeroNode) InsertNode(head *HeroNode, newNode *HeroNode) {
	//1.循环找到最后一个节点插入 
	tmpNode := head
	for {
		if tmpNode.next == nil { 
			break
		}
		tmpNode = tmpNode.next //找到最后一个节点插入
	}
	tmpNode.next = newNode
}

//显示链表
func (this *HeroNode) ListNode(head *HeroNode) {
	//1.判断该链表是否为空
	tmpNode := head
	if tmpNode.next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		fmt.Printf("[%d %s %s]--->",tmpNode.next.no,tmpNode.next.name,tmpNode.next.nickname)
		tmpNode = tmpNode.next
		if tmpNode.next == nil {
			break
		}
	}
}

//按编号从小到大插入链表
func  (this *HeroNode) InsertNodeByNo(head *HeroNode, newNode *HeroNode) (err error) {
	tmpNode := head
	flag := true
	for {
		if tmpNode.next == nil { //到链表最后跳出
			break
		} else if tmpNode.next.no > newNode.no {
			break
		} else if tmpNode.next.no == newNode.no { //链表中存在就不差入
			flag = false
			break
		}
		tmpNode = tmpNode.next
	}
	if !flag { 
		errInfo := fmt.Sprintf("链表中以存在该节点:%v", newNode)
		return errors.New(errInfo)
	} else { 
		newNode.next = tmpNode.next
		tmpNode.next = newNode
	}
	return
}

func main()  {
	head := &HeroNode{}
	hero1 := &HeroNode{
		no : 1,
		name : "宋江",
		nickname : "及时雨",
	}
	hero2 := &HeroNode{
		no : 2,
		name : "卢俊义",
		nickname : "玉麒麟",
	}
	hero3 := &HeroNode{
		no : 3,
		name : "林冲",
		nickname : "豹子头",
	}
	hero4 := &HeroNode{
		no : 4,
		name : "吴用",
		nickname : "智多星",
	}
	// head.InsertNode(head, hero1)
	// head.InsertNode(head, hero3)
	// head.InsertNode(head, hero4)
	// head.InsertNode(head, hero2)
	head.InsertNodeByNo(head, hero1)
	head.InsertNodeByNo(head, hero3)
	head.InsertNodeByNo(head, hero4)
	head.InsertNodeByNo(head, hero2)
	head.ListNode(head)
}