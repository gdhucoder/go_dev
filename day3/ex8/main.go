package main

import "fmt"

// 写一个程序，获取一个变量的地址，并打印到终端
func printAdd(a *int) {
	fmt.Println(a)
}

func changeVal(a *int) {
	*a = 2 * *a
}

func main() {
	a := 10
	printAdd(&a)
	changeVal(&a)
	fmt.Println(a)

	// 指针类型
	var p *int
	p = &a
	fmt.Println(*p)

	*p = 100
	fmt.Println(a)

}
