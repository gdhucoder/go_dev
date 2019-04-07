package main

import (
	"fmt"
)

func main() {
	n()
	m()
	n()
}

var a string = "GO"

func n() {
	fmt.Println(a)
}

func m() {
	// a := "O" // 声明，并初始化一个局部变量a。 和全局变量a不一样。
	// a = "O" // 修改全局变量的值
	fmt.Println(a)
}
