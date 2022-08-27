package session

import (
	"fmt"
	"net"
	"sync"
	"testing"
)

func TestSessionReadWrite(t *testing.T) {
	//定义监听端口
	addr := "127.0.0.1:8000"
	//定义发送数据
	myData := "hello"
	//定义协程组
	wg := sync.WaitGroup{}
	wg.Add(2)
	//写数据协程
	go func() {
		defer wg.Done()
		//创建tcp链接
		listen, err := net.Listen("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}
		s := NewSession(conn)
		//写数据
		err = s.Write([]byte(myData))
		if err != nil {
			t.Fatal(err)
		}
	}()

	//读数据协程
	go func() {
		defer wg.Done()
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			t.Fatal(err)
		}
		s := NewSession(conn)
		data, err := s.Read()
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != myData {
			fmt.Println("数据读取错误")
		}
		fmt.Println(string(data))
	}()

	wg.Wait()
	fmt.Println("程序执行结束")
}
