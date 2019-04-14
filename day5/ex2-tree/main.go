package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	left  *Student
	right *Student
}

func traverse(p *Student) {

	if p == nil {
		return
	}

	traverse(p.left)
	fmt.Println(p)
	traverse(p.right)

	// if p.left != nil {
	// 	traverse(p.left)
	// }

	// fmt.Println(p)
	// if p.right != nil {
	// 	traverse(p.right)
	// }

}

func main() {
	var root *Student = &Student{
		Name:  "wtw",
		Age:   1,
		Score: 100,
	}

	var stu1 = &Student{
		Name:  "1",
		Age:   1,
		Score: 100,
	}

	var stu2 = &Student{
		Name:  "2",
		Age:   1,
		Score: 100,
	}

	root.left = stu1
	root.right = stu2

	var stu3 = &Student{
		Name:  "3",
		Age:   1,
		Score: 100,
	}

	var stu4 = &Student{
		Name:  "4",
		Age:   1,
		Score: 100,
	}

	stu1.left = stu3
	stu1.right = stu4

	traverse(root)

}
