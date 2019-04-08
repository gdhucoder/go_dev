package main

import (
	"fmt"
	"strings"
)

const pre = "https"

func main() {
	str1 := "https://xxxxxxxxxxxxxxxxx"
	str2 := "xxxx"
	fmt.Println(strings.HasPrefix(str1, pre)) // true
	fmt.Println(strings.HasPrefix(str2, pre)) // false

	// =============================================
	
}
