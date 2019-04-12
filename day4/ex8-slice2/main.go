package main

import "fmt"

func test() {
	var a [5]int = [...]int{1, 3, 5, 7, 9}
	b := a[1:]
	fmt.Printf("s = %p, a[1] = %p \n", b, &a[1])

	b[0] = 999
	fmt.Println(a) // 由于b[0]和a[1]的地址相同，所以改变b[0]的值，a[1]也跟着改变了
	// [1 999 5 7 9]

	b = append(b, 10) // 又重新拷贝了一份， 根据容量的大小重新分配内存
	fmt.Printf("s = %p, a[1] = %p \n", b, &a[1])
}

// 两个切片append，使用...
func testAppend() {
	var a = []int{123, 1234, 123, 123, 123}
	var b = []int{45, 45, 4, 545, 45}
	a = append(a, b...) // 两个切片append
	fmt.Println(a)
}

func testCopy() {

	var a = []int{123, 234, 34, 34, 34, 34, 3434}
	b := make([]int, 1)
	copy(b, a)
	fmt.Println(b)

}

// 关于创建切片
func testMake() {
	a := make([]int, 2, 10) // 创建切片
	fmt.Printf("len %d, capacity %d\n", len(a), cap(a))
	// len 2, capacity 10
}

func main() {
	testMake() // 观察 make 的使用 len 和 cap 参数
	testCopy()
	testAppend()
	test()
}
