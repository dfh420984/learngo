package session

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
	"time"
)

//用于测试的机构体
type User struct {
	Name string
	Age  int
}

//查询用户的方法
func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	user[0] = User{Name: "zs", Age: 20}
	user[1] = User{Name: "ls", Age: 21}
	user[2] = User{Name: "ww", Age: 22}
	//模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("uid %d not in user db", uid)
}

//测试方法
func TestRpc(t *testing.T) {
	//需要对interface{}可能产生的类型进行注册
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	//创建服务端
	srv := NewServer(addr)
	//将方法注册到服务端
	srv.Register("queryUser", queryUser)
	//服务端等待调用
	go srv.Run()
	time.Sleep(time.Millisecond)
	//客户端链接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	//创建客户端
	cli := NewClient(conn)
	//声明一个函数原型
	var query func(int) (User, error)
	cli.callRPC("queryUser", &query)
	//得到查询结果
	u, err := query(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(u)
}
