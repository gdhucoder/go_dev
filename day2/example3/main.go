package main

import (
	"fmt"
	"time"
)

// 定义常量的写法
const (
	Man    = 1
	Female = 2
)

// 第一个数自动初始化为0， 后面的数字初始为 1 ，2 ，3
const (
	ZERO = iota
	ONE
	SECOND
	THIRD
)

func main() {
	fmt.Println("third: ", THIRD)
	for {
		// 获取时间的秒数
		second := time.Now().Unix()
		fmt.Println(second)
		if second%Female == 0 {
			fmt.Println("female")
		} else {
			fmt.Println("male")
		}
		time.Sleep(1000 * time.Millisecond)
	}

}
