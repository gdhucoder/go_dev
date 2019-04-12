package main

import "fmt"

func test() {
	//
	var a [5]int = [5]int{5, 4, 3, 2, 1}
	fmt.Println(a)

	// 类型推导
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 省略个数
	var c = [...]int{1, 3, 5}
	fmt.Println(c)

	// 指定初始哪个元素
	var d = [...]int{1: 100, 3: 200}
	fmt.Println(d)
	var e = [56]int{1: 100, 3: 200}
	fmt.Println(e)

}

// 二维数组的初始化和遍历
func testMulArr() {
	var a = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(a)

	var b = [...][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(b)

	for i, v := range a {
		for j, v1 := range v {
			fmt.Printf("(%d ,%d)-->%d ", i, j, v1)
		}
		fmt.Println()
	}

}

func main() {

	// 数组的各种初始化
	test()
	// 二维数组的初始化和遍历
	testMulArr()
}
