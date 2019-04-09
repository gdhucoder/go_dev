package main

import "fmt"

// 函数返回值可以取个名字
func concat(str string, arg ...string) (res string) {
	res = str
	for i := 0; i < len(arg); i++ {
		res += arg[i]
	}
	return
}

func main() {
	res := concat("hello", " ", "world!")
	fmt.Println(res)
	res = concat("你好", "中国", "学习强国！")
	fmt.Println(res)
}
