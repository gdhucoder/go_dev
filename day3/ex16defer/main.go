package main

import "fmt"

func main() {
	// defer 推迟
	var a int = 0
	defer fmt.Println("a", a)
	a = 10
	defer fmt.Println("a", a)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
