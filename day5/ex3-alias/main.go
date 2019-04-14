package main

import "fmt"

type integer int

func main() {
	var i integer = 1000
	var j = int(i)
	fmt.Println(j)
	// int 和 integer是不同的类型了，相互使用需要强制类型转换
}
