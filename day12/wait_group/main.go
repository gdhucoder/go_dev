package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	// wg.Add(10)
	for i := 0; i < 10; i++ {
		wg.Add(1) // recommand put code here
		go calc(&wg, i)
		// wg.Done() // not ok
	}

	wg.Wait()
	fmt.Println("all goroutine finish")

	// calc: 3
	// calc: 2
	// calc: 7
	// calc: 6
	// calc: 4
	// calc: 9
	// calc: 8
	// calc: 0
	// calc: 1
	// all goroutine finish

}
func calc(w *sync.WaitGroup, i int) {

	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done() // finish a goroutine
}
