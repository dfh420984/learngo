package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标识位置
	mutexWoken                   // 唤醒标识位置
	mutexStarving                // 锁饥饿标识位置
	mutexWaiterShift = iota      // 标识waiter的起始bit位置
)

// 扩展一个Mutex结构，拓展额外功能
type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (mu *Mutex) TryLock() bool {
	// 没有其他线程争抢，如果能成功抢到锁,&mu.Mutex 代表结构体首地址，指向state
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&mu.Mutex)), 0, mutexLocked) {
		return true
	}
	// 如果处于唤醒、加锁或者饥饿状态，这次请求就不参与竞争了，返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&mu.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}
	// 尝试在竞争的状态下请求锁
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&mu.Mutex)), old, new)
}

func (mu *Mutex) try() {
	go func() {
		mu.Lock()
		num := rand.Intn(2)
		time.Sleep(time.Duration(num) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	ok := mu.TryLock() // 尝试获取到锁
	if ok {
		fmt.Println("got the lock") // do something mu.Unlock() return
		mu.Unlock()
		return
	}
	// 没有获取到
	fmt.Println("can't get the lock")
}

//当前持有和等待这把锁的 goroutine 的总数
func (mu *Mutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&mu.Mutex)))
	v = v >> mutexWaiterShift //得到等待者的数值
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}

// 锁是否被持有
func (mu *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&mu.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 是否有等待者被唤醒
func (mu *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&mu.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (mu *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&mu.Mutex)))
	return state&mutexStarving == mutexStarving
}

func count() {
	var mu Mutex
	for i := 0; i < 1000; i++ { // 启动1000个goroutine
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	// 输出锁的信息
	fmt.Printf("waitings: %d, isLocked: %t, woken: %t,  starving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

//使用 Mutex 实现一个线程安全的队列
type SliceQueue struct {
	data []interface{}
	mu sync.Mutex
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, n),
	}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func main() {
	// var mu Mutex
	// mu.try()
	//count()
}
