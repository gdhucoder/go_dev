package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("今天是几号")
	today := time.Now().Weekday()
	fmt.Printf("%T\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Too far away")
	}
}
