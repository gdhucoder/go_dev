package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (p *Student) SetName(name string) *Student {
	p.Name = name
	return p
}

func (p *Student) SetAge(age int) *Student {
	p.Age = age
	return p
}

func main() {

	stu := &Student{}
	// link pattern operations
	stu.SetName("hgd").SetAge(32)

	fmt.Println(stu)

}
