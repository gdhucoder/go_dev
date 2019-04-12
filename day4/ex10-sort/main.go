package main

import (
	"fmt"
	"sort"
)

// 整数排序
func testSortInt() {
	var a = [...]int{61, 77, 1, 2, 999, 99} // 值类型
	sort.Ints(a[:])                         // 切片，变成引用类型
	fmt.Println(a)
}

// 字符串类型排序
func testSortString() {
	var s = [...]string{"23", "aaaab", "ab", "A", "Axxaa"}
	sort.Strings(s[:])
	fmt.Println(s)
}

//
func testSortFloat() {
	var a = [...]float64{1.2, 1.3, 0.1, 0.12}
	sort.Float64s(a[:])
	fmt.Println(a)
	// [0.1 0.12 1.2 1.3]
}

func testIntSearch() {
	// 前提数组时有序的
	var a = [...]int{3, 2, 1}
	sort.Ints(a[:])
	idx := sort.SearchInts(a[:], 1)
	fmt.Println(idx)
}

func testStringSearch() {
	// TODO
}

func testFloatSearch() {
	// TODO
}

func main() {
	testIntSearch() // 搜索
	testSortFloat()
	testSortString()
	testSortInt()
	// [23 A Axxaa aaaab ab]
	// [1 2 61 77 99 999]
}
