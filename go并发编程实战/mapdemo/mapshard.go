package main

import (
	"sync"
)

//减少锁的粒度常用的方法就是分片（Shard），将一把锁分成几把锁，每个锁控制一个分片，以降低锁的粒度，进而提高访问此 map 对象的吞吐。Go 比较知名的分片并发 map 的实现是
//https://github.com/orcaman/concurrent-map
//扩展其它功能的 map 实现，比如带有过期功能的timedmap、使用红黑树实现的 key 有序的treemap等
//https://github.com/zekroTJA/timedmap
//https://pkg.go.dev/github.com/emirpasic/gods/maps/treemap

//分片加锁：更高效的并发 map
//分片数量
var SHARD_COUNT = 32

// 分成SHARD_COUNT个分片的map
type ConcurrentMap []*ConcurrentMapShared

// 通过RWMutex保护的线程安全的分片，包含一个map
type ConcurrentMapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

// 创建并发map
func New() ConcurrentMap {
	m := make(ConcurrentMap, SHARD_COUNT)
	for i := 0; i < SHARD_COUNT; i++ {
		m[i] = &ConcurrentMapShared{
			items: make(map[string]interface{}),
		}
	}
	return m
}

// 根据key计算分片索引
func (m ConcurrentMap) GetShard(key string) *ConcurrentMapShared {
	return m[uint(fnv32(key))%uint(SHARD_COUNT)]
}

// FNV hash
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func (m ConcurrentMap) Set(key string, value interface{}) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.Lock() //对这个分片加锁，执行业务操作
	shard.items[key] = value
	shard.Unlock()
}

func (m ConcurrentMap) Get(key string) (interface{}, bool) {
	// 根据key计算出对应的分片
	shard := m.GetShard(key)
	shard.RLock()
	val, ok := shard.items[key]
	shard.Unlock()
	return val, ok
}
