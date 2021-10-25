package innerfunc

import "fmt"

func NewMap() {
	m := new(map[int]string)
	fmt.Printf("new map:%v", m)
}
