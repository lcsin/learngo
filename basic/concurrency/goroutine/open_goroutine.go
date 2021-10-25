/*
1. 使用go关键字创建goroutine,一个goroutine对应一个函数
2. main函数也是一个goroutine,程序启动时会默认为main函数创建一个goroutine
3. main函数的goroutine执行结束返回后，其创建的所有子goroutine都会结束
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)
}
