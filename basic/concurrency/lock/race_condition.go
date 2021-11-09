/*
Race Condition:
	1. 竞争条件:一个系统或进程的输出依赖于不受控制的事件出现顺序或出现时机,例如几个线程对同一个共享资源的写入和读取产生了重叠

counter++
	1. 不是一个原子操作,无法在一个cpu时间内计算完成
	2. 其操作有以下三步:
		i:先从内存中读取counter的值
	   ii:计算+1后counter的值
      iii:将+1后的值再赋值给counter
	3. 因为counter++不是原子操作，所以再自增过程中存在被其他goroutine复写的可能

sync.Mutex:
	1. go语言基础包自带的互斥锁
	2. mtx.Lock(): 加锁,解锁之前只能一个goroutine进入
	3. mtx.Unlock(): 解锁

sync.atomic:
	1. 封装了原子操作,是线程安全的
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg1     sync.WaitGroup
	mtx     sync.Mutex
	counter int32
	ch      = make(chan int, 1)
)

func UnSafeIncCounter() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		counter++ // 共享资源
	}
}

func MutexIncCounter() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()   // 加锁
		counter++    // 共享资源
		mtx.Unlock() // 解锁
	}
}

func AtomicIncCounter() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt32(&counter, 1)
	}
}

func ChannelIncCounter() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		count := <-ch
		count++
		ch <- count
	}
}

func main() {
	wg1.Add(2)

	go ChannelIncCounter()
	go ChannelIncCounter()
	ch <- 0

	wg1.Wait()
	fmt.Println("task over! counter:", <-ch)
}
