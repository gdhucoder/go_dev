package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(a, b string) string {
	// 123
	//   8
	// res
	res := ""
	isAddOne := false
	for i := len(a) - 1; i >= 0; i-- {
		num := (a[i] - '0') + (b[i] - '0')
		if isAddOne {
			num += 1
		}
		if num >= 10 {
			isAddOne = true
			if i == 0 {
				res = "" + string(num/10+'0') + string(num-10+'0') + res
			} else {
				res = "" + string(num-10+'0') + res
			}
		} else {
			res = "" + string(num+'0') + res
			isAddOne = false
		}
	}
	// fmt.Println(res)
	return res
}

func equalLength(a, b *string) {
	diff := len(*a) - len(*b)
	if diff > 0 {
		*b = strings.Repeat("0", diff) + *b
	} else {
		*a = strings.Repeat("0", -diff) + *a
	}
}
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
	a := "123"
	b := "99"
	equalLength(&a, &b)
	fmt.Println("a", a)
	fmt.Println("b", b)
	c := add(a, b)
	fmt.Println("c", c)
	aint, _ := strconv.Atoi(a)
	bint, _ := strconv.Atoi(b)
	sum := aint + bint
	fmt.Println("sum", sum)

}
