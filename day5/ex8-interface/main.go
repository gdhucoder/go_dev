package main

import "fmt"

type Test interface {
	Print()
	Sleep()
}

type Student struct {
	Name string
	Age  int
}

func (p *Student) Print() {
	fmt.Println("student: ", p)
}

func (p *Student) Sleep() {
	fmt.Println("student sleep")
}

type People struct {
	Name string
	Age  int
}

func (p *People) Print() {
	fmt.Println("people: ", p)
}

func (p *People) Sleep() {
	fmt.Println("people sleep")
}

func main() {
	var t Test
	t.Print() // 接口就是一个地址（指针）
	var stu Student
	stu.Name = "huguodong"
	stu.Age = 33

	t = &stu // 接口可以指向这个类型
	t.Print()
	t.Sleep()

	var p People = People{
		Name: "wtw",
		Age:  1,
	}

	t = &p
	t.Print()
	t.Sleep()

	// student:  &{huguodong 33}
	// student sleep
	// people:  &{wtw 1}
	// people sleep

}
