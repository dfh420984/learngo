package session

import (
	"net"
	"reflect"
)

//声明客户端
type Client struct {
	conn net.Conn
}

//创建客户端对象
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

//实现一个通用的rpc客户端
//绑定rpc访问的方法
//传入访问的函数名

//函数具体实现实在server端，客户端只有函数原型
//使用MakeFunc完成原型到函数的调用

//fPtr 指向函数原型
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	//通过反射获取fPtr为初始化的原型
	fn := reflect.ValueOf(fPtr).Elem()
	//另一个函数，作用是对第一个函数参数操作
	//完成于server的交互
	f := func(args []reflect.Value) []reflect.Value {
		//处理输入的参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		//创建链接
		cliSession := NewSession(c.conn)
		//编码数据
		reqRpc := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRpc)
		if err != nil {
			panic(err)
		}
		//写出数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		//读取想应数据
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		//解码数据
		respRpc, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		//处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRpc.Args))
		for i, arg := range respRpc.Args {
			//必须进行nil转换
			if arg == nil {
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}

	//参数1，为初始化函数的方法值，类型是relect.Type
	//参数2，另一个函数，作用是对第一个函数参数操作
	//返回relect.Value类型
	//MakeFunc 使用传入的函数原型， 创建一个绑定参数2的新函数
	v := reflect.MakeFunc(fn.Type(), f)
	//为函数的fPtr赋值
	fn.Set(v)
}
