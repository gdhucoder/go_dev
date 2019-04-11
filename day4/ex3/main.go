package main

import "fmt"

func recur(n int) {
	fmt.Println(n)
	recur(n + 1)
}

func main() {
	recur(0)
}
