package main

import "fmt"

func send(intChan chan int, exitChan chan bool) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)
	exitChan <- true
}

func recevice(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok { // ok = false 管道关闭了
			break
		}
		fmt.Println("received : ", v)
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 2)
	go send(intChan, exitChan)
	go recevice(intChan, exitChan)

	var total int
	for _ = range exitChan {
		total++ // 一个计数
		fmt.Printf("total : %d\n", total)
		if total == 2 {
			break
		}
	}
}
