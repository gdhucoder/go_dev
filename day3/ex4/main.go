package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 100)
}

func main() {
	now := time.Now()
	// 固定，这个格式，这个是go语言诞生的时间
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	// 统计程序的执行耗时
	start := time.Now().UnixNano()
	test()
	end := time.Now().UnixNano()
	fmt.Printf("cost: %d\n", (end-start)/1000)
}
