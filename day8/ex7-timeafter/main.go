package main

import (
	"fmt"
	"time"
)

func main() {
	// 定时器
	t := time.After(time.Second)
	fmt.Println("time.After", <-t)

}
