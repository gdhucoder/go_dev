package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func traverse(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head *Student = &Student{
		Name:  "wtw",
		Age:   1,
		Score: 100,
	}

	// 声明针类型
	var stu1 *Student = new(Student)
	stu1.Name = "hgd"
	stu1.Age = 22
	stu1.Score = 99

	head.next = stu1

	traverse(head)

	var stu2 Student
	stu2.Name = "huy"
	stu1.next = &stu2

	traverse(head)

	// &{wtw 1 100 0xc00005e300}
	// &{hgd 22 99 <nil>}

	// 	{wtw 1 100 0xc00006a2d0}
	// {hgd 22 99 <nil>}

}
