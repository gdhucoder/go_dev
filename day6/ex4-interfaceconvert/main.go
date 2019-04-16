package main

import "fmt"

type Student struct {
	Name string
}

func Test(a interface{}) {
	b, ok := a.(int)
	if ok == false {
		fmt.Println("Convert failed!")
		return
	}
	b += 3
	fmt.Println(b)
}

func classifier(items ...interface{}) {
	for i, v := range items {
		// fmt.Println(v.(type)) // 不能这样用
		switch v.(type) {
		case bool:
			fmt.Printf("%d param is bool, value is %v\n", i, v)
		case int, int32, int64:
			fmt.Printf("%d param is int, value is %v\n", i, v)
		case float32, float64:
			fmt.Printf("%d param is float, value is %v\n", i, v)
		case string:
			fmt.Printf("%d param is string, value is %v\n", i, v)
		case Student:
			fmt.Printf("%d param is student, value is %v\n", i, v)
		case *Student:
			fmt.Printf("%d param is student ptr, value is %v\n", i, v)
		}
	}
}

func main() {
	var a interface{}
	var b int
	a = b
	c := a.(int)
	fmt.Println(c)
	Test(c)
	var s string
	Test(s)

	var stu Student = Student{
		Name: "wtw",
	}

	var stu1 = &Student{
		Name: "hgd",
	}

	classifier(28, 1.1, "Hello", true, stu, stu1)
	// 0 param is int, value is 28
	// 1 param is float, value is 1.1
	// 2 param is string, value is Hello
	// 3 param is bool, value is true
	// 4 param is student, value is {wtw}
	// 5 param is student ptr, value is &{hgd}
}
