package main

import "fmt"

// deferred function calls are pushed onto a stack
// when a function returns, its deferred calls are executed
// in last-in-first-out order
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
	// counting
	// done
	// 9
	// 8
	// 7
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
	// 0
}
