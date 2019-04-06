package main

import (
	"fmt"
)

func modify(a int) {
	a = 10
}

// 使用指针类型，传入a的地址
func modify1(a *int) {
	*a = 10 // ** 表示指向的值
}

func main() {
	a := 5
	b := make(chan int, 1) // 引用类型
	fmt.Println("a=", a)
	// 输出的b为地址
	fmt.Println("b=", b)

	modify(a)
	fmt.Println("a=", a)

	modify1(&a) // a 的地址
	fmt.Println("a=", a)

	// a= 5
	// b= 0xc0000120e0
	// a= 5
	// a= 10
}
