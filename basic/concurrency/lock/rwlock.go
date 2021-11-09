package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	count int
	wg    sync.WaitGroup
	lock  sync.RWMutex
)

func write() {
	defer wg.Done()
	lock.Lock() // 加写锁
	count++
	time.Sleep(10 * time.Millisecond)
	lock.Unlock() // 解写锁
}

func read() {
	defer wg.Done()
	lock.RLock() // 加读锁
	time.Sleep(time.Millisecond)
	lock.RUnlock() // 解读锁
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
