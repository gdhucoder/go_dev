package main

import "fmt"

// add 一个多个整数相加
func add(a int, arg ...int) (sum int) {
	sum = a
	// if len(arg) >= 1 {
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	// }
	return
}

func main() {
	sum := add(1, 2, 3)
	fmt.Println(sum)
	sum = add(1)
	fmt.Println(sum)
	sum = add(1, 3, 7, 10)
	fmt.Println(sum)
}
