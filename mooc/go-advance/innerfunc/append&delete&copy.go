package innerfunc

import "fmt"

func AppendElement() {
	slice := make([]string, 4)
	slice[0] = "张三"
	slice[1] = "李四"
	fmt.Printf("len:%d, cap:%d, val:%v\n", len(slice), cap(slice), slice)
	newSlice := append(slice, "王五")
	fmt.Printf("len:%d, cap:%d, val:%v\n", len(newSlice), cap(newSlice), newSlice)
}

func CopyElement() {
	slice1 := make([]string, 3)
	slice1[0] = "张三"
	slice1[1] = "李四"
	slice1[2] = "田七"
	fmt.Println("before copy slice1:", slice1)

	slice2 := make([]string, 3)
	slice2[0] = "王五"
	slice2[1] = "赵六"
	fmt.Println("before copy slice2:", slice2)

	copy(slice1, slice2)
	fmt.Println("after copy slice1:", slice1)
	fmt.Println("after copy slice2:", slice2)
}

func DeleteMap() {
	m := make(map[int]string)
	m[0] = "张三"
	m[1] = "李四"
	m[2] = "王五"
	fmt.Println("before delete:", m)
	delete(m, 1)
	fmt.Println("after delete:", m)
}
