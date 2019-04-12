package main

import "fmt"

// string 底层是byte数组，可以进行slice操作
func testString() {
	str := "hhh eee"
	s1 := str[0:3]
	s2 := str[4:]
	fmt.Println(s1)
	fmt.Println(s2)
}

func testModifyString() {

	// 英文字符
	s := "hello world"
	//	s[1] = 'E' // 字符串是不可变的
	s1 := []byte(s) // 转换成byte数组，是可变的了
	s1[0] = 'X'
	str := string(s1)
	fmt.Println(str)

	// 中文字符
	s = "我在深圳，hello"
	s2 := []rune(s)
	s2[0] = '你'
	s = string(s2)
	fmt.Println(s)

}

func main() {
	testModifyString()
	testString()
}
