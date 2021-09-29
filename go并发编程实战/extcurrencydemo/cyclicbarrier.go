package main

import (
	"context"
	"log"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/marusama/cyclicbarrier"
	"golang.org/x/sync/semaphore"
)

// 并发趣题：一氧化二氢制造工厂
// 有一个名叫大自然的搬运工的工厂，生产一种叫做一氧化二氢的神秘液体。这种液体的分子是由一个氧原子和两个氢原子组成的，也就是水。这个工厂有多条生产线，每条生产线负责生产氧原子或者是氢原子，每条生产线由一个 goroutine 负责。这些生产线会通过一个栅栏，只有一个氧原子生产线和两个氢原子生产线都准备好，才能生成出一个水分子，否则所有的生产线都会处于等待状态。也就是说，一个水分子必须由三个不同的生产线提供原子，而且水分子是一个一个按照顺序产生的，每生产一个水分子，就会打印出 HHO、HOH、OHH 三种形式的其中一种。HHH、OOH、OHO、HOO、OOO 都是不允许的。生产线中氢原子的生产线为 2N 条，氧原子的生产线为 N 条。
// 1。首先，我们来定义一个 H2O 辅助数据类型，它包含两个信号量的字段和一个循环栅栏。semaH 信号量：控制氢原子。一个水分子需要两个氢原子，所以，氢原子的空槽数资源数设置为
// 2。semaO 信号量：控制氧原子。一个水分子需要一个氧原子，所以资源数的空槽数设置为 1。
// 3。循环栅栏：等待两个氢原子和一个氧原子填补空槽，直到任务完成。

// 定义水分子合成的辅助数据结构
type H2O struct {
	semaH *semaphore.Weighted         // 氢原子的信号量
	semaO *semaphore.Weighted         //氧原子的信号量
	b     cyclicbarrier.CyclicBarrier //循环栅栏，用来控制合成
}

func New() *H2O {
	return &H2O{
		semaH: semaphore.NewWeighted(2), //氢原子需要两个
		semaO: semaphore.NewWeighted(1), // 氧原子需要一个
		b:     cyclicbarrier.New(3),     // 需要三个原子才能合成
	}
}

// 流水线分为氢原子处理流水线和氧原子处理流水线，首先，我们先看一下氢原子的流水线：如果有可用的空槽，氢原子的流水线的处理方法是 hydrogen，hydrogen 方法就会占用一个空槽（h2o.semaH.Acquire），输出一个 H 字符，然后等待栅栏放行。等其它的 goroutine 填补了氢原子的另一个空槽和氧原子的空槽之后，程序才可以继续进行。
func (h2o *H2O) hydrogen(releaseHydrogen func()) {
	h2o.semaH.Acquire(context.Background(), 1)
	releaseHydrogen()                 // 输出H
	h2o.b.Await(context.Background()) //等待栅栏放行
	h2o.semaH.Release(1)              //释放氢原子空槽
}

//然后是氧原子的流水线。氧原子的流水线处理方法是 oxygen， oxygen 方法是等待氧原子的空槽，然后输出一个 O，就等待栅栏放行。放行后，释放氧原子空槽位。
func (h2o *H2O) oxygen(releaseOxygen func()) {
	h2o.semaO.Acquire(context.Background(), 1)
	releaseOxygen()                   // 输出O
	h2o.b.Await(context.Background()) //等待栅栏放行
	h2o.semaO.Release(1)              // 释放氢原子空槽
}

//在栅栏放行之前，只有两个氢原子的空槽位和一个氧原子的空槽位。只有等栅栏放行之后，这些空槽位才会被释放。栅栏放行，就意味着一个水分子组成成功
func TestWaterFactory() {
	//用来存放水分子结果的channel
	var ch chan string
	releaseHydrogen := func() { ch <- "H" }
	releaseOxygen := func() { ch <- "O" }
	// 300个原子，300个goroutine,每个goroutine并发的产生一个原子
	var N = 100
	ch = make(chan string, N*3)
	h2o := New()
	// 用来等待所有的goroutine完成
	var wg sync.WaitGroup
	wg.Add(N * 3)
	// 200个氢原子goroutine
	for i := 0; i < 2*N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.hydrogen(releaseHydrogen)
			wg.Done()
		}()
	}
	// 100个氧原子goroutine
	for i := 0; i < N; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			h2o.oxygen(releaseOxygen)
			wg.Done()
		}()
	}
	//等待所有的goroutine执行完
	wg.Wait()
	// 结果中肯定是300个原子
	if len(ch) != N*3 {
		log.Printf("expect %d atom but got %d \n", N*3, len(ch))
	}
	// 每三个原子一组，分别进行检查。要求这一组原子中必须包含两个氢原子和一个氧原子，这样才能正确组成一个水分子
	var s = make([]string, 3)
	for i := 0; i < N; i++ {
		s[0] = <-ch
		s[1] = <-ch
		s[2] = <-ch
		sort.Strings(s)
		water := s[0] + s[1] + s[2]
		if water != "HHO" {
			log.Printf("expect a water molecule but got %s \n", water)
		} else {
			log.Printf("get a water %s \n", water)
		}
	}
}

func main() {
	TestWaterFactory()
}
