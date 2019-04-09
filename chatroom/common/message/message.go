package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
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

type RegisterResMes struct {
	Code int `json:"code"`  //返回状态码 400表示该用户已存在，200表示注册成功
	Error string `json:"error"` //返回错误信息
}

