package main

import "fmt"

func statChar(str string) (numOfCN, numOfEN, numOfNum, numOfSpace, numOfOtherChar int) {
	s := []rune(str)
	// var (
	// 	numOfCN, numOfEN, numOfNum, numOfSpace, numOfOtherChar = 0, 0, 0, 0, 0
	// )
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= 19968 && s[i] <= 40896:
			numOfCN++
		case s[i] >= 48 && s[i] <= 57:
			numOfNum++
		case (s[i] >= 97 && s[i] <= 122) || (s[i] >= 65 && s[i] <= 90):
			numOfEN++
		case s[i] == rune(' '):
			numOfSpace++
		default:
			numOfOtherChar++
		}
	}
	return
}

// 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
func main() {

	str := `你好世界 hello world 666 \n\r\t`
	// 	unicode编码范围：
	// 汉字：[0x4e00,0x9fa5]（或十进制[19968,40869]）
	// 数字：[0x30,0x39]（或十进制[48, 57]）
	// 小写字母：[0x61,0x7a]（或十进制[97, 122]）
	// 大写字母：[0x41,0x5a]（或十进制[65, 90]）
	// 空格：十进制32
	fmt.Println(statChar(str))

	fmt.Println(rune(' '))

}
