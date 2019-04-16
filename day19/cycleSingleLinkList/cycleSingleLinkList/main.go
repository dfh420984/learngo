package main

import (
	"fmt"
)

//单向循环链表
type HeroNode struct {
	no int
	name string
	nickname string
	next *HeroNode
}

//插入链表
func (this *HeroNode) InsertNode(head *HeroNode, newNode *HeroNode) {
	//1.判断是不是第一个添加
	if head.next == nil {
		head.no = newNode.no
		head.name = newNode.name
		head.nickname = newNode.nickname
		head.next = head  //自己指向自己
		fmt.Println(newNode,"加入到了链表中")
		return
	}
	//2.找到环形的最后一个节点
	tmpNode := head
	for {
		if tmpNode.next == head { 
			break
		}
		tmpNode = tmpNode.next //找到最后一个节点插入
	}
	tmpNode.next = newNode
	newNode.next = head
}

//输出单向环形链表
func (this *HeroNode) ListNode(head *HeroNode) {
	//1.判断该链表是否为空
	tmpNode := head
	if tmpNode.next == nil {
		fmt.Println("该链表为空")
		return
	}
	for {
		fmt.Printf("[%d %s %s]--->",tmpNode.no,tmpNode.name,tmpNode.nickname)
		if tmpNode.next == head {
			break
		}
		tmpNode = tmpNode.next
	}
}

func (this *HeroNode) DelNode(head *HeroNode, id int) *HeroNode {
	tmpNode := head
	helper := head
	if tmpNode.next == nil {
		fmt.Println("这是一个空的环形链表")
		return head
	}
	//只有一个节点
	if tmpNode.next == head { 
		if tmpNode.no == id {
			tmpNode.next = nil
		}
		return head
	}
	//将helper定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	//如果有两个以上节点
	flag := true
	for {
		if tmpNode.next == head { //比较到最后一个，最后一个还没比较
			break
		}
		if tmpNode.no == id {
			if tmpNode == head { //说明删除的是头节点
				head = head.next
			}
			helper.next = tmpNode.next
			flag = false
			break
		}
		tmpNode = tmpNode.next
		helper = helper.next
	}
	//这里还有比较一次
	if flag { //如果flag 为真，则我们上面没有删除
		if tmpNode.no == id {
			helper.next = tmpNode.next
		}else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	} 
	return head
}

//单向循环列表
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
	head.InsertNode(head, hero1)
	head.InsertNode(head, hero2)
	head.InsertNode(head, hero3)
	head.InsertNode(head, hero4)
	head = head.DelNode(head, 3)
	head.ListNode(head)
}