package main

import "fmt"

func main() {
	i := 0
LABEL: // 使用goto 实现for 循环
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto LABEL
}
