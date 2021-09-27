package main

import "fmt"

//生成一个数据流
func asStream(done <-chan struct{}) <-chan interface{} {
	s := make(chan interface{})
	values := []int{1, 2, 3, 4, 5}
	go func() {
		defer close(s)
		for _, v := range values {
			select {
			case <-done: //退出
				return
			case s <- v:
			}
		}
	}()
	return s
}

func mapChan(in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{}) //创建一个输出chan
	if in == nil {
		close(out)
		return out
	}
	go func() { //启动一个goroutine,实现map的主要逻辑
		defer close(out)
		for v := range in {
			out <- fn(v)
		}
	}()
	return out
}

func reduce(in <-chan interface{}, fn func(r, v interface{}) interface{}) interface{} {
	if in == nil {
		return nil
	}
	out := <-in // 先读取第一个元素
	for v := range in {
		out = fn(out, v)
	}
	return out
}

func main() {
	done := make(chan struct{})
	in := asStream(done)
	// map操作: 乘以10
	mapFn := func(v interface{}) interface{} {
		return v.(int) * 10
	}
	// reduce操作: 对map的结果进行累加
	reduceFn := func(r, v interface{}) interface{} {
		return r.(int) + v.(int)
	}
	sum := reduce(mapChan(in, mapFn), reduceFn) //返回累加结果
	close(done)
	fmt.Println(sum)
}
