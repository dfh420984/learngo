package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

//这里我们定义几个用户状态的常量
const (
	UserOnline = iota
	UserOffline 
	UserBusyStatus 
)

type Message struct {
	Type string `json:"type"`  //消息类型
	Data string `json:"data"`  //消息数据
}

//登陆结构体
type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`  //返回状态码 500表示该用户未注册，200表示登陆成功
	Error string `json:"error"` //返回错误信息
	UsersId []int //增加字段用来保存用户id切片
}

//注册结构体
type RegisterMes struct {
	User User `json:"user"`
}

//注册返回消息结构体
type RegisterResMes struct {
	Code int `json:"code"`  //返回状态码 400表示该用户已存在，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态
}

//群发结构体
type SmsMes struct {
	Content string `json:"content"`
	User //匿名结构体，继承
}

