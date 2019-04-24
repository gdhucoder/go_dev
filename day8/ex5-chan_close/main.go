package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)

	for {
		b, ok := <-intChan
		if !ok {
			break // 如果不关闭 死循环
		}
		fmt.Println(b)
	}

	// 	0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
}
