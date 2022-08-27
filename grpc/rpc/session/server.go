package session

import (
	"fmt"
	"net"
	"reflect"
)

//声明服务端
type Server struct {
	//地址
	addr string
	//服务端函数名到函数反射值映射map
	funcs map[string]reflect.Value
}

//创建服务端对象
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

//服务端绑定注册方法
//将函数名和函数实现真正对应起来
//第一个参数为函数名，第二个参数为真正函数
func (s *Server) Register(rpcName string, f interface{}) {
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	//map中没有，将映射添加到map中， 便于调用
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

//服务端等待调用
func (s *Server) Run() {
	//监听
	listen, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("监听 %s, err: %v", s.addr, err)
		return
	}
	for {
		//拿到链接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		//创建会话
		serSession := NewSession(conn)
		//rpc读取数据
		b, err := serSession.Read()
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		//对数据解码
		rpcData, err := decode(b)
		if err != nil {
			fmt.Println("decode err", err)
			return
		}
		//根据读取到的数据，获得调用的函数名
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("函数 %s 不存在\n", rpcData.Name)
			return
		}
		//解析遍历客户端传递过来的参数
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		//反射调用方法，传入参数
		out := f.Call(inArgs)
		//解析遍历执行结果
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		//包装数据，返回给客户端
		respRpcData := RPCData{rpcData.Name, outArgs}
		respBytes, err := encode(respRpcData)
		if err != nil {
			fmt.Println("encode err", err)
			return
		}
		//使用rpc写出数据
		err = serSession.Write(respBytes)
		if err != nil {
			fmt.Println("session write err", err)
			return
		}
	}
}
