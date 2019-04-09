package main

import "fmt"

// 自定义类型，函数是一种数据类型
type op_func func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func operator(op op_func, a, b int) int {
	return op(a, b)
}

func main() {
	c := add
	sum := c(10, 20)
	fmt.Println(sum)
	// if type(c) == add {
	// 	fmt.Println("c equals to add func")
	// }
	sum = operator(add, 100, 102)

	fmt.Println(sum)

	sub := operator(sub, 102, 1)
	fmt.Println(sub)

	// 30
	// 202
	// 101

	fmt.Println(c)
	fmt.Println(add)

}
