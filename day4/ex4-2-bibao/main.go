package main

import (
	"fmt"
	"strings"
)

func divider(n int) func(int) int {
	return func(a int) int {
		return a / n
	}
}

// 后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	func1 := makeSuffixFunc(".bmp")
	fmt.Println(func1("test"))
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func2("test"))

	half := divider(2)

	fmt.Println(half(30))

}
