package main

import "sync"

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
	delete(m.m , k)
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