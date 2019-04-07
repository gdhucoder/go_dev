package main

import (
	"fmt"
	"strconv"
)

func main() {
	// str := "abcdefghijklmnopqrstuvwxyz"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%d\n", str[i])
	// }
	var str1 string
	var result int = 0
	fmt.Scanf("%s", &str1)
	fmt.Println(str1)
	for i := 0; i < len(str1); i++ {
		num := int(str1[i] - '0')
		result += (num * num * num)
	}
	// 将字符串转换成数字
	number, _ := strconv.Atoi(str1)
	if number == result {
		fmt.Println("是水仙花数", result)
	} else {
		fmt.Println("不是水仙花数")
	}

}
