package main

import (
	"fmt"
	"strings"
)

func main() {
	// 计算两个大数相加的和，这两个大数会超过int64的表示范围
	str1 := "100000999349349"
	fmt.Println(str1[1] - '0')
	str2 := "2"
	var str3 string
	if len(str1) > len(str2) {
		str3 = strings.Repeat("0", len(str1)-len(str2)) + str2
	}
	fmt.Println(str1)
	fmt.Println(str3)

}
