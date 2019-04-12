package main

import "fmt"

// 函数中数组为值传递！！！！
func change(arr [5]int) {
	arr[0] = 100
	fmt.Println("inchange: ", arr)
}

// 如果要修改数组的值，要使用地址
func test3(arr *[5]int) {
	(*arr)[0] = 1000
	arr[2] = 999 // 这样也对
	a := arr     //
	a[4] = 999999999999
	fmt.Println(&a)
	fmt.Println(arr)

	// 为什么？
	// 	&[1000 0 0 0 999999999999]
	// &[1000 0 0 0 0]
}

func main() {
	var a [5]int
	a[0] = 1
	// j := 10
	// a[j] = 100
	fmt.Println(a)

	// 遍历数组的方式：一
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	// 使用i,v := range a
	for i, v := range a {
		fmt.Println(i, v)
	}

	b := a
	b[1] = 100
	fmt.Println(a)
	fmt.Println(b)

	change(a)
	fmt.Println("out change ", a)

	// var c [6]int
	// change(c) // 参数中数组的长度是5， 传入6就报错了，因为是两种类型

	// 函数中修改arr的值
	test3(&a)
	fmt.Println(a)
}
