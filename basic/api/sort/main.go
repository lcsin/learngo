/*
@see https://learnku.com/articles/38269
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

type Student struct {
	Name string
	Age  int
}

type sortAge []Student

func (s sortAge) Len() int {
	return len(s)
}

func (s sortAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s sortAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	sortWithInt()
	sortWithStructAge()
	sortWithStructAgeName()
	sortWithInterface()
}

func sortWithInt() {
	arr := []int{4, 2, 3, 1}
	fmt.Println("排序前：", arr)

	sort.Ints(arr)
	fmt.Println("排序后：", arr)
}

func sortWithStructAge() {
	student := []Student{
		{"张三", 23},
		{"李四", 2},
		{"王五", 2},
		{"赵六", 25},
	}
	fmt.Println("排序前：", student)

	sort.SliceStable(student, func(i, j int) bool {
		return student[i].Age < student[j].Age
	})
	fmt.Println("排序后：", student)
}

// order by age asc, name desc
func sortWithStructAgeName() {
	student := []Student{
		{"张三", 23},
		{"李四", 2},
		{"王五", 2},
		{"赵六", 25},
	}
	fmt.Println("排序前：", student)

	sort.SliceStable(student, func(i, j int) bool {
		if student[i].Age != student[j].Age {
			return student[i].Age < student[j].Age
		}
		return strings.Compare(student[i].Name, student[j].Name) == 1
	})

	fmt.Println("排序后：", student)
}

func sortWithInterface() {
	student := []Student{
		{"张三", 23},
		{"李四", 2},
		{"王五", 2},
		{"赵六", 25},
	}
	fmt.Println("排序前：", student)

	sort.Sort(sortAge(student))

	fmt.Println("排序后：", student)
}
