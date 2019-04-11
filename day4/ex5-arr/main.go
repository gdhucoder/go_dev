package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 1
	j := 10
	a[j] = 100
	fmt.Println(a)
}
