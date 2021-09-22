package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

//获取runtime.Stack 方法获取栈帧信息，栈帧信息里包含 goroutine id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)

	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

//Mutex 不是可重入的锁。
func main() {
	l := &sync.Mutex{}
	foo(l)
}
