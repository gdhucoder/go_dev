package main

import (
	"fmt"
	// 包可以起别名
	a "go_dev/day2/example2/add"
)

func main() {
	a.Test()
	fmt.Println("您的名字是", a.Name)
	fmt.Println("您的年龄是", a.Age)
}
