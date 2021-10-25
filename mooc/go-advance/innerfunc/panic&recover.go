package innerfunc

import "fmt"

func Panic() {
	defer coverPanic()
	panic("error argument")
}

func coverPanic() {
	msg := recover()
	switch msg.(type) {
	case string:
		fmt.Println("string msg:", msg)
	case error:
		fmt.Println("error msg:", msg)
	default:
		fmt.Println("unknown msg:", msg)
	}
}
