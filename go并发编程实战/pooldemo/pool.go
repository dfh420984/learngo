package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

//使用 sync.Pool 做 buffer 池，但是，如果用刚刚的那种方式做 buffer 池的话，可能会有内存泄漏的风险。取出来的 bytes.Buffer 在使用的时候，我们可以往这个元素中增加大量的 byte 数据，这会导致底层的 byte slice 的容量可能会变得很大。这个时候，即使 Reset 再放回到池子中，这些 byte slice 的容量不会改变，所占的空间依然很大。而且，因为 Pool 回收的机制，这些大的 Buffer 可能不被回收，而是会一直占用很大的空间，这属于内存泄漏的问题。在使用 sync.Pool 回收 buffer 的时候，一定要检查回收的对象的大小，太大就丢弃不回收。

var bufPool = sync.Pool{
	New: func() interface{} {
		// 初始化，通常返回指针类型
		return new(bytes.Buffer)
	},
}

func timeNow() time.Time {
	return time.Unix(time.Now().Unix(), 0)
}

func Log(w io.Writer, key, val string) {
	//从bufPool中取走一个元素，这也就意味着，这个元素会从 Pool 中移除，返回给调用者
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	//将buffer中的数据输出
	w.Write(b.Bytes())
	//将一个元素返还给 Pool，Pool 会把这个元素保存到池中
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
}
