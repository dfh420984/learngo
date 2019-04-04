package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"reflect"
	_"time"
)

func redis_set(conn redis.Conn) {
	_, err := conn.Do("SET", "name", "dfh420984")
	if err != nil {
		fmt.Println("redis set error", err)
		return
	}
	_, err = conn.Do("expire", "name", 10) //10秒过期
    if err != nil {
        fmt.Println("set expire error: ", err)
        return
    }
	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error", err)
		return
	}
	fmt.Println("redis get name", name)

	//redis批量设置mset mget
	_, err = conn.Do("MSET", "name", "wd","age",22)
    if err != nil {
        fmt.Println("redis mset error:", err)
    }
    res, err := redis.Strings(conn.Do("MGET", "name","age"))
    if err != nil {
        fmt.Println("redis get error:", err)
    } else {
        res_type := reflect.TypeOf(res)
        fmt.Printf("res type : %s \n", res_type)
        fmt.Printf("MGET name: %s \n", res)
        fmt.Println(len(res))
    }
}

func redis_list(conn redis.Conn) {
	_, err := conn.Do("LPUSH", "list1", "ele1","ele2","ele3")
    if err != nil {
        fmt.Println("redis mset error:", err)
    }
    res, err := redis.String(conn.Do("LPOP", "list1"))
    if err != nil {
        fmt.Println("redis POP error:", err)
    } else {
        res_type := reflect.TypeOf(res)
        fmt.Printf("res type : %s \n", res_type)
        fmt.Printf("res  : %s \n", res)
    }
}

func redis_hash(conn redis.Conn) {
	_, err := conn.Do("HSET", "student","name", "wd","age",22)
    if err != nil {
        fmt.Println("redis mset error:", err)
    }
    res, err := redis.Int64(conn.Do("HGET", "student","age"))
    if err != nil {
        fmt.Println("redis HGET error:", err)
    } else {
        res_type := reflect.TypeOf(res)
        fmt.Printf("res type : %s \n", res_type)
        fmt.Printf("res  : %d \n", res)
    }

}

func redis_pipeline(conn redis.Conn) {
	conn.Send("HSET", "student","name", "wd","age","22")
    conn.Send("HSET", "student","Score","100")
    conn.Send("HGET", "student","age")
    conn.Flush()

    res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
    res2, err := conn.Receive()
    fmt.Printf("Receive res2:%v\n",res2)
    res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n",res3)
	fmt.Printf(" res err:%v \n", err)

}

func redis_subs() { 
	conn,err := redis.Dial("tcp","0.0.0.0:6382")
	if err != nil {
		fmt.Println("redis sub 连接错误", err)
		return
	}
	psc := redis.PubSubConn{conn}
    psc.Subscribe("channel1") //订阅channel1频道
    for {
        switch v := psc.Receive().(type) {
        case redis.Message:
            fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
        case redis.Subscription:
            fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
        case error:
            fmt.Println(v)
            return
        }
    }
}

func redis_pub(message string) { 
	conn,err := redis.Dial("tcp","0.0.0.0:6382")
	if err != nil {
		fmt.Println("redis pub 连接错误", err)
		return
	}
	_,err1 := conn.Do("PUBLISH", "channel1", message)
	if err1 != nil {
			fmt.Println("pub err: ", err1)
			return
	}
}

func redis_trans() {
	conn,err := redis.Dial("tcp","0.0.0.0:6382")
	if err != nil {
		fmt.Println("redis 连接错误", err)
		return
	}
	defer conn.Close()
    conn.Send("MULTI")
    conn.Send("INCR", "foo")
    conn.Send("INCR", "bar")
    r, err := conn.Do("EXEC")
    fmt.Println(r)

}

var Pool redis.Pool
func init()  {      //init 用于初始化一些参数，先于main执行
    Pool = redis.Pool{
			MaxIdle:     16,
			MaxActive:   32,
			IdleTimeout: 120,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "0.0.0.0:6382")
        },
    }
}


func main()  {
	// conn,err := redis.Dial("tcp","0.0.0.0:6382")
	// if err != nil {
	// 	fmt.Println("redis连接错误", err)
	// 	return
	// }
	// fmt.Println("redis连接成功")
	// defer conn.Close()
	//redis_set(conn)
	//redis_list(conn)
	//redis_hash(conn)
	//redis_pipeline(conn)
	// go redis_subs()
	// go redis_pub("this is wd")
	// time.Sleep(time.Second*3)
	//redis_trans()
	conn := Pool.Get()
    res,err := conn.Do("HSET","student","name","jack")
    fmt.Println(res,err)
    res1,err := redis.String(conn.Do("HGET","student","name"))
    fmt.Printf("res:%s,error:%v",res1,err)
}