package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {
	conn,err := redis.Dial("tcp","0.0.0.0:6382")
	if err != nil {
		fmt.Println("redis连接错误", err)
	}
	fmt.Println("redis连接成功", conn)
	
}