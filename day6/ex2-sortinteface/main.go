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

type ByAge []Student

func (s ByAge) Len() int {
	return len(s)
}

func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

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
	fmt.Println(stus)
}
