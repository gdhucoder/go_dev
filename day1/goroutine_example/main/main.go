package main

import (
	"fmt"
	"go_dev/day1/goroutine_example/goroutine"
)

func main() {
	c := make(chan int, 1)
	go goroutine.Add(10, 20, c)
	num := <-c // 省略声明
	fmt.Println(num)
}
