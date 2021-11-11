package main

import (
	"fmt"
	"math"
)

func round(x float64) int {
	return int(math.Floor(x + 0/5))
}

type User struct {
	Name string
	Age  int
}

func main() {

	//fmt.Printf("%.2f", float64(4)/float64(52)*100)

	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%T\n", i)
	//}

	var user User
	fmt.Println(user)
	fmt.Printf("%v,%T\n", user, user)
	fmt.Println(&user == nil)
}
