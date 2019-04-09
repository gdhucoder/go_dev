package main

import "fmt"

// 匿名函数
var (
	result = func(a, b int) int {
		return a + b
	}
)

func test(a, b int) int {
	result := func(a1, b1 int) int {
		return a1 + b1
	}(a, b)
	// return result(a, b)
	return result
}

func main() {
	res := test(100, 300)
	fmt.Println(res)
	res = result(1, 2)
	fmt.Println(res)
}
