package main

import (
	"fmt"
)

/*
交换两个数的值
*/
func swap(a *int, b *int) {
	temp := *a // 取 a 地址对应的内存的值
	*a = *b
	*b = temp
}

func swap1(a int, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("swap")
	a := 10
	b := 5
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// 第二种方法

	a, b = swap1(a, b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	// 第三种方法
	a, b = b, a
	fmt.Println("a=", a)
	fmt.Println("b=", b)

}
