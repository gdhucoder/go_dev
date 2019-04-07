package main

import (
	"fmt"
)

func main() {
	var a int8 = 100
	var b int16
	b = int16(a) // 强制类型转换
	fmt.Printf("a=%d b=%d \n", a, b)

	var n int16 = 34
	var m int32
	// m = n
	m = int32(n)
	fmt.Println(m)
	fmt.Println(n)
}
