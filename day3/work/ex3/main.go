package main

import (
	"fmt"
)

func isHuiWen(str string) bool {

	// 变量字符串类型使用 rune，不论中英文，就是一个字符
	s := []rune(str)
	fmt.Println(s)
	fmt.Println(len(s))

	for i := 0; i <= (len(s)-1)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "上海自来水来自海上"
	if isHuiWen(str) {
		fmt.Println(str)
	}
}
