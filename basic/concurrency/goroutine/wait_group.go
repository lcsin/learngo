/*
1. sync.WaitGroup用于实现goroutine之间的同步(等待所有并发任务完成)
2. wg.Add(): 每启动一个goroutine计数+1,相当于有几个任务在执行
3. wg.Wait(): 等待所有计数的goroutine结束,等待所有任务完成
4. wg.Done(): 每当一个goroutine结束计数-1,表示任务完成
*/
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("goroutine:", i)
}

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println("all goroutine over...")
}
