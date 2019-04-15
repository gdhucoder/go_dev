package main

import "fmt"

type integer int

func (p integer) print() {
	fmt.Println(p)
}

func (p *integer) set(b integer) {
	*p = b
}

type Student struct {
	Name string
	Age  int
}

// 这么写会传递p的一个副本
// 结构体中的方法 中间包含一个（类型）
func (p Student) init(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println(p)
}

func (p *Student) init2(name string, age int) {
	p.Name = name
	p.Age = age
	fmt.Println(p)
}

func (p Student) get() {
	fmt.Println(p.Age, p.Name)
}

func main() {
	var stu Student
	// (&stu).init2("huguodong", 33)
	stu.init2("huguodong", 33) // 省略了 & go人性化
	// stu.init("huguodong", 33)
	fmt.Println("out", stu)

	var a integer
	a = 100
	a.print()
	// &{huguodong 33}
	// out {huguodong 33}
	// 100

	a.set(10000)
	a.print()
	// 10000
}
