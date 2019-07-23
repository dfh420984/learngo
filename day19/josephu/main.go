package main

import (
	"fmt"
)

type Boy struct {
	no   int
	next *Boy
}

/*********josephu问题************/

//添加小孩
func addBoy(num int) *Boy {
	//1.初始化
	first := &Boy{}
	cur := &Boy{}
	if num < 1 {
		fmt.Println("num参输不合法")
		return first
	}
	//2.开始添加
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 { //添加第一个小孩
			first = boy
			cur = boy
			cur.next = first
		} else {
			cur.next = boy
			cur = boy
			cur.next = first //构造环形链表
		}
	}
	return first
}

//显示单向循环链表
func showLinkList(first *Boy) {
	if first.next == nil {
		fmt.Println("该单向循环链表为空")
		return
	}
	cur := first
	for {
		fmt.Printf("小孩标号：%d \n", cur.no)
		if cur.next == first {
			break
		}
		cur = cur.next
	}
}

//统计链表节点个数
func countList(first *Boy) int {
	count := 0
	if first.next == nil {
		return count
	}
	cur := first
	for {
		count++
		if cur.next == first {
			break
		}
		cur = cur.next
	}
	return count
}

//开始josephu
func playGame(first *Boy, startNo int, countNo int) {
	if first.next == nil {
		fmt.Println("该单向循环链表为空,没有小孩")
		return
	}
	totalNum := countList(first)
	if startNo > totalNum {
		fmt.Println("开始编号小孩过大")
		return
	}
	//定义辅助指针，帮助删除小孩节点
	tail := first
	//让tail指向循环链表最后一个节点
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}
	//开始从第startNo个小孩,
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	for {
		//开始数countNo，first指针直到对应小孩
		for i := 1; i <= countNo-1; i++ {
			first = first.next
			tail = tail.next
		}
		//开始删除小孩
		fmt.Printf("小孩编号%d出圈\n", first.no)
		first = first.next
		tail.next = first
		if tail == first { //最后剩下一个小孩
			break
		}
	}
	fmt.Printf("小孩编号%d出圈\n", first.no)
}

func main() {
	first := addBoy(5)
	showLinkList(first)
	// count := countList(first)
	// fmt.Println("链表中元素个数为：",count)
	playGame(first, 2, 3)
}
