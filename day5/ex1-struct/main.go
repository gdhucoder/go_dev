package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	score float32
}

// 结构体里面所有字段内存布局都是连续的
// 例如stu的内存地址以Name的内存地址开头
func main() {
	var stu Student
	stu.Name = "wangtiawaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	stu.Age = 1
	stu.score = 99
	fmt.Println(stu)
	fmt.Printf("stu %p \n", &stu)
	fmt.Printf("Name: %p\n", &stu.Name)
	fmt.Printf("Age %p\n", &stu.Age)
	fmt.Printf("score: %p\n", &stu.score)

	var stu1 *Student = &Student{
		Name: "胡国栋",
		Age:  3,
	}

	fmt.Println(stu1)
	fmt.Println(stu1.Name) // 可以省略 * 析取符号

	var stu2 = Student{
		Name: "董杰出",
	}
	fmt.Println(stu2)
	// 	{wangtiawaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa 1 99}
	// stu 0xc00004e400
	// Name: 0xc00004e400
	// Age 0xc00004e410
	// score: 0xc00004e418
	// &{胡国栋 3 0}
	// 胡国栋
	// {董杰出 0 0}
}
