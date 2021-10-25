package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//innerfunc.MakeSlice()
	//innerfunc.MakeMap()
	//innerfunc.MakeChannel()

	//innerfunc.NewMap()

	//innerfunc.AppendElement()
	//innerfunc.CopyElement()
	//innerfunc.DeleteMap()

	//innerfunc.Panic()

	//json.SerializeMap()
	//json.SerializeStruct()
	//json.UnSerializeStruct()
	//json.SerializeMap()

	str := "hello,world!"
	fmt.Println(len(str))
	str = "你好，世界!"
	fmt.Println(utf8.RuneCountInString(str))

	ch := make(chan int, 3)
	fmt.Println(cap(ch))
}
