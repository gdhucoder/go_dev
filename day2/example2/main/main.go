package main

import (
	"fmt"
	// 包可以起别名
	a "go_dev/day2/example2/add"
)

// 执行main之前，默认执行init函数
func init() {
	fmt.Println("in main initialized")
}

func main() {
	a.Test()
	fmt.Println("您的名字是", a.Name)
	fmt.Println("您的年龄是", a.Age)
}
