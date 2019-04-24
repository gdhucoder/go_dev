package main

import (
	"fmt"
	"time"
)

// chan 为指针类型
func write(intChan chan int) {
	for i := 0; i < 100; i++ {
		intChan <- i
		fmt.Printf("writing %d ...\n", i)
	}
}

func read(intChan chan int) {
	for {
		i := <-intChan
		fmt.Println(i)
	}
}

func main() {

	// 数字大小为缓冲区
	intChan := make(chan int, 5)
	go write(intChan)
	go read(intChan)
	time.Sleep(time.Second * 4)

}
