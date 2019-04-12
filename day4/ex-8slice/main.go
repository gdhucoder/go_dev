package main

import "fmt"

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

func main() {
	testSlice()
}
