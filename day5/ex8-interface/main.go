package main

import "fmt"

type Test interface {
	Print()
}

type Student struct {
	Name string
	Age  int
}

func (p *Student) Print() {
	fmt.Println(p)
}

func main() {
	var stu Student
	stu.Name = "huguodong"
	stu.Age = 33

	var t Test
	t = &stu // 接口可以指向这个类型
	t.Print()
}
