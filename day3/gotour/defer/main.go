package main

import "fmt"

// defer statement defers the execution
// of a function until the surrounding function returns
func main() {
	// 推迟，先执行周围的函数
	defer fmt.Println("world")
	fmt.Println("hello")
}
