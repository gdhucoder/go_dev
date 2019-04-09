package main

import (
	"fmt"
	"strings"
)

func Print(n int) {
	var str string
	for i := 1; i <= n; i++ {
		str = strings.Repeat("A", i)
		fmt.Println(str)
	}
}

func main() {
	Print(3)
	str := "hello world 中国"
	for i, v := range str {
		fmt.Printf(
			"index[%d] val[%c] type %T len[%d] \n",
			i, v, v, len([]byte(string(v))))
	}
}
