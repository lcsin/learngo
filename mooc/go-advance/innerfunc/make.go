package innerfunc

import "fmt"

// MakeSlice 通过make函数创建slice
func MakeSlice() {
	strings := make([]string, 3)
	strings[0] = "dog"
	strings[1] = "cat"
	strings[2] = "duck"
	fmt.Printf("make slice:%v\n", strings)
}

// MakeMap 通过make函数创建map
func MakeMap() {
	m := make(map[string]interface{}, 3)
	m["code"] = 200
	m["msg"] = "请求成功"
	m["data"] = nil
	fmt.Printf("make map:%v\n", m)
}

// MakeChannel 通过make函数创建channel
func MakeChannel() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 2
	ch <- 1
	fmt.Printf("make channel:%v\n", ch)
}
