/*
make(chan int):
	1. 创建一个没有缓存的channel
	2. 直接往没有缓存的channel中塞值时,会引发阻塞
value, ok := <-ch:
	1. 从chan中获取值,返回值为value, bool
	2. 当chan被关闭时,bool返回false
	3. 当chan中暂时没有值时,unbuffered chan会阻塞,直到chan中有值
ch <- value:
	1. 把值存储到chan中
	2. 当unbuffered chan中有值,但是没有被goroutine接收时,程序会处于阻塞状态
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

func player(name string, ch chan int) {
	defer wg.Done()

	for {
		ball, ok := <-ch
		// 通道关闭
		if !ok {
			fmt.Printf("channel is closed! %s wins!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%10 == 0 {
			close(ch)
			fmt.Printf("%s miss the ball! %s loses!\n", name, name)
			return
		}

		ball++
		fmt.Printf("%s receive the ball:%d\n", name, ball)
		ch <- ball
	}
}

func main() {
	wg.Add(2)

	ch := make(chan int)

	go player("张三", ch)
	go player("李四", ch)

	ch <- 0

	wg.Wait()
	fmt.Println("task over...")
}
