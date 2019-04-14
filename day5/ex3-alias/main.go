package main

import "fmt"

type integer int

type Stu Student

type Student struct {
	Number int
}

func main() {
	var i integer = 1000
	var j = int(i)
	fmt.Println(j)
	// int 和 integer是不同的类型了，相互使用需要强制类型转换

	var stu1 Student = Student{30}

	var stu2 Stu
	stu2 = Stu(stu1) // 强转
	fmt.Println(stu2)
}
