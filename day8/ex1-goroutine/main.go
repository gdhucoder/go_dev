package main

import (
	"fmt"
	"time"
)

func test() {
	var i int
	for {
		fmt.Println(i)
		i++
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for {
		fmt.Println("running in main")
		time.Sleep(time.Second * 2)
	}
	// 0
	// running in main
	// 1
	// 2
	// running in main
	// 3
	// 4
	// running in main
}
