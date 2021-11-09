/*
运行结果是有序同步的
*/
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			v, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- v * v
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
	fmt.Println("finished...")
}
