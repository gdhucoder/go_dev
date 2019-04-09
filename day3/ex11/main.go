package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 猜数字的游戏
func main() {
	var n int
	rand.Seed(time.Now().UnixNano())
	n = rand.Intn(101)
	// fmt.Println("generate random number: ", n)
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		switch {
		case input < n:
			fmt.Println("less")
		case input > n:
			fmt.Println("greater")
		default:
			fmt.Println("equal")
			return
		}
	}

}
