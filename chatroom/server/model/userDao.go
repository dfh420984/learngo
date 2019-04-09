package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
)

//定义一个全局userDao
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//工厂模式返回userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool : pool,
	}
	return 
}

//根据一个用户id，返回一个User实例
func (this *UserDao) getUserById(id int) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("HGET","users", id))
	if err != nil { 
		if err == redis.ErrNil { //表示在redis hash中,没有找到对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//这里把res反序列成User实例
	err = json.Unmarshal([]byte(res), user) 
	if err != nil {
		fmt.Println("json.Unmarshal user  err=", err)
		return 
	}
	return
}

//完成登录校验
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	//1.先判断id是否存在
	user, err = this.getUserById(userId)
	if err != nil {
		return
	}
	//2.再来判断密码是否正确 
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}