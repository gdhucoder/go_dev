package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("name: %s, age: %d\n", s.Name, s.Age)
}

// 学生按照年龄排序， 实现排序接口的三个函数
type ByAge []Student

// 长度
func (s ByAge) Len() int {
	return len(s)
}

// 小于
func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// 交换
func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {

	var stus ByAge

	for i := 0; i < 10; i++ {
		var stu = Student{
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		stus = append(stus, stu)
	}

	sort.Sort(stus)
	fmt.Println(stus) // 已经按照年龄排好序
	// name: stu56, age: 0
	// name: stu37, age: 6
	// name: stu94, age: 11
	// name: stu81, age: 18
	// name: stu25, age: 40
	// name: stu11, age: 45
	// name: stu47, age: 59
	// name: stu28, age: 74
	// name: stu81, age: 87
	// name: stu62, age: 89

}
