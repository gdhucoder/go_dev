package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic ", err)
		}
	}()
	var m map[string]int
	m["age"] = 1
}

func calc() {
	for {
		fmt.Println("calc...")
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	for i := 0; i < 2; i++ {
		go calc()
	}
	time.Sleep(time.Second * 10)
	// 出现了panic程序继续运行
	// calc...
	// calc...
	// panic  assignment to entry in nil map
	// calc...
	// calc...
}
