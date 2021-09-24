package main

import (
	"sync"
)

//在 Go 中，map[key]函数返回结果可以是一个值，也可以是两个值
// func main() {
//     var m = make(map[string]int)
//     m["a"] = 0
//     fmt.Printf("a=%d; b=%d\n", m["a"], m["b"])

//     av, aexisted := m["a"]
//     bv, bexisted := m["b"]
//     fmt.Printf("a=%d, existed: %t; b=%d, existed: %t\n", av, aexisted, bv, bexisted)
// }

//常见错误一：未初始化
// func main() {
//     var m map[int]int //解决办法初始化这个实例（m := make(map[int]int)
//     m[100] = 100
// }

//常见错误二：并发读写
//虽然这段代码看起来是读写 goroutine 各自操作不同的元素，貌似 map 也没有扩容的问题，但是运行时检测到同时对 map 对象有并发访问，就会直接 panic
// func main() {
//     var m = make(map[int]int,10) // 初始化一个map
//     go func() {
//         for {
//             m[1] = 1 //设置key
//         }
//     }()

//     go func() {
//         for {
//             _ = m[2] //访问这个map
//         }
//     }()
//     select {}
// }

//加读写锁：扩展 map，支持并发读写
type RWMap struct {
	sync.RWMutex
	m map[int]int
}

// 新建一个RWMap
func NewRWMap(n int) *RWMap {
	return &RWMap{m: make(map[int]int, n)}
}

//从map中读取一个值
func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k] // 在锁的保护下从map中读取
	return v, existed
}

// 设置一个键值对
func (m *RWMap) Set(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = k
}

// 删除一个键值对
func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

// map的长度
func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

// 遍历map
func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock() //遍历期间一直持有读锁
	defer m.RUnlock()
	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}

//Go 内建的 map 类型不是线程安全的，所以 Go 1.9 中增加了一个线程安全的 map，也就是 sync.Map。但是，我们一定要记住，这个 sync.Map 并不是用来替换内建的 map 类型的，它只能被应用在一些特殊的场景里。在以下两个场景中使用 sync.Map，会比使用 map+RWMutex 的方式，性能要好得多：
// 1.只会增长的缓存系统中，一个 key 只写入一次而被读很多次；
// 2.多个 goroutine 为不相交的键集读、写和重写键值对。
// func syncMapDemo() {

// 	var smp sync.Map

// 	// 数据写入
// 	smp.Store("name", "小红")
// 	smp.Store("age", 18)

// 	// 数据读取
// 	name, _ := smp.Load("name")
// 	fmt.Println(name)

// 	age, _ := smp.Load("age")
// 	fmt.Println(age)

// 	// 遍历
// 	smp.Range(func(key, value interface{}) bool {
// 		fmt.Println(key, value)
// 		return true
// 	})

// 	// 删除
// 	smp.Delete("age")
// 	age, ok := smp.Load("age")
// 	fmt.Println("删除后的查询")
// 	fmt.Println(age, ok)

// 	// 读取或写入,存在就读取，不存在就写入
// 	smp.LoadOrStore("age", 100)
// 	age, _ = smp.Load("age")
// 	fmt.Println("不存在")
// 	fmt.Println(age)

// 	smp.LoadOrStore("age", 99)
// 	age, _ = smp.Load("age")
// 	fmt.Println("存在")
// 	fmt.Println(age)
// }
