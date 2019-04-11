package main

import "fmt"

func test() {
	s1 := new([]int) // 切片
	fmt.Println(s1)
	// (*s1)[0] = 100 // 会报错，因为没有申请空间
	*s1 = make([]int, 5)
	fmt.Println(s1)
	// &[0 0 0 0 0]
	s2 := make([]int, 10)
	fmt.Println(s2)
	// &[]
	// [0 0 0 0 0 0 0 0 0 0]
	return
}

func main() {
	test()
}
