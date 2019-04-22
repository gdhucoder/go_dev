package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		// recover会使程序从panic中恢复，并返回panic value
		if e := recover(); e != nil {
			fmt.Printf("panicking %s\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\n")
}

func main() {
	fmt.Printf("Calling test\n")
	test()
	fmt.Printf("Test complete\n")
}
