package main

import (
	"fmt"
)

func calc(intChan, resChan chan int, exitChan chan bool) {
	for v := range intChan {
		var isPrime bool
		isPrime = true
		for i := 2; i < v/2; i++ {
			if v%i == 0 {
				isPrime = false
			}
		}
		if isPrime {
			resChan <- v
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 10)
	resChan := make(chan int, 10)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 1000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resChan, exitChan)
	}

	// 等待所有计算的goroute全部退出
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
			fmt.Printf("%d exited...\n", i)
		}
		close(resChan) // 关闭 resChan
		fmt.Println("resChan closed!")
	}()

	// fatal error: all goroutines are asleep - deadlock!
	// for i := 0; i < 8; i++ {
	// 	<-exitChan
	// }
	// close(resChan)
	for v := range resChan {
		fmt.Println(v)
	}

	// 977
	// 0 exited...
	// 1 exited...
	// 2 exited...
	// 3 exited...
	// 4 exited...
	// 5 exited...
	// 997
	// 6 exited...
	// 7 exited...
	// resChan closed!

}
