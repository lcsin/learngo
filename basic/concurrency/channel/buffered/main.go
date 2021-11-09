/*
make(chan int, 5):
	1. 创建5个缓存大小的channel
	2. 往有缓存的channel中的发送操作只有在channel满的时候才会阻塞
	3. 有缓存的通道关闭后,依然可以从通道中获取值
for range可以遍历一个通道,在遍历通道获取值前需要close(chan)

*/
package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	close(ch)
	for val := range ch {
		fmt.Println(val)
	}

}
