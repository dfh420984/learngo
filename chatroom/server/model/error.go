package model

import (
	"errors"
)

//自定义一些错误信息
var (
	ERROR_USER_NOTEXISTS = errors.New("用户不存在..")
	ERROR_USER_EXISTS = errors.New("用户已经存在..")
	ERROR_USER_PWD = errors.New("用户密码不正确..")
)