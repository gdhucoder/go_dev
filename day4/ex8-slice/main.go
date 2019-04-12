package main

import "fmt"

// 切片的底层数据结构
type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 9999
}

func testSelfDefineSlice() {
	var s1 slice
	s1 = make1(s1, 10)
	modify(s1)
	fmt.Println(s1.ptr)
}

func testSlice() {
	var a [5]int = [...]int{1, 3, 4, 5435, 234}
	var slice []int
	slice = a[1:]
	fmt.Println(slice)
	fmt.Println(len(slice))
	// 0<=len<=cap
	fmt.Println(cap(slice))
	slice = slice[0:1]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func testSlice4() {

	var a [4]int = [4]int{1, 2, 3, 4}
	b := a[1:] // 切片指向数组第一个元素的指针

	// 地址是一样的
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])

}

func main() {
	testSlice4()
	testSelfDefineSlice()
	testSlice()
}
