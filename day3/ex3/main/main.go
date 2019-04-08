package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	test3() // index
	test4() // Lastindex

	begin, end := StrIndex("abbbbbbbbbb", "b")
	fmt.Printf("begin = %d, end = %d\n", begin, end)

	test5()

	test6()

	test7()

	test8() // to lower

	test9()

	test10()

	test11()
}

func test3() {
	str1 := "aaaaabbbccc"
	pos := strings.Index(str1, "b")
	fmt.Println("出现的位置： ", pos)
}

func test4() {
	str1 := "abccccc"
	pos := strings.Index(str1, "c")
	fmt.Println("last index: ", pos)
}

func test5() {
	str1 := "abbbbbbbb"
	str2 := strings.Replace(str1, "bb", "c", 3)
	fmt.Println(str2)
}

func test6() {
	str1 := "abbb"
	count := strings.Count(str1, "b")
	fmt.Printf("count = %d\n", count)
}

func test7() {
	str1 := "abbb"
	str2 := strings.Repeat(str1, 2)
	fmt.Printf("str2 = %s\n", str2)
}

func test8() {
	str1 := "AAAAAAAAAAAAAAAA"
	str2 := strings.ToLower(str1)
	str3 := strings.ToUpper(str2)
	fmt.Println(str2)
	fmt.Println(str3)
}

func test9() {
	str1 := "   aa   "
	str2 := strings.TrimSpace(str1) // 	去掉首位空白符
	fmt.Println(str2)
}

func test10() {
	fmt.Println("==================test10")
	str1 := "bbbbbaabbb"
	str2 := strings.Trim(str1, "ba") //
	fmt.Println("Trim", str2)
	str2 = strings.TrimLeft(str1, "b") //
	fmt.Println("Trim", str2)
	str2 = strings.TrimRight(str1, "bbb") //
	fmt.Println("Trim", str2)
}

func test11() {
	fmt.Println("==================test11")
	str1 := "aa bb cc"
	str2 := strings.Fields(str1)                   // 返回str空格分隔的所有子串的slice
	str3 := strings.Split(str1, " ")               // 返回str split 分隔的所有子串的slice
	str4 := strings.Join(str2, "----------------") // 用sep把s1中的所有元素连接起来
	fmt.Println("Fields", str2[1:])
	fmt.Println("Split", str3[0])
	fmt.Println("Str4 Join", str4)
	str5 := strconv.Itoa(1111) // 数字转化成字符串
	fmt.Println("strconv.Itoa", str5)
	number, err := strconv.Atoi("aaa")
	if err != nil {
		fmt.Println("can not conver to int", err)
		return
	}
	fmt.Println("strconv.Atoi: ", number)

}

// 写一个函数返回一个字符串在另一个字符串的首次出现和最后出现的位置

func StrIndex(str string, substr string) (int, int) {
	start := strings.Index(str, substr)
	end := strings.LastIndex(str, substr)
	return start, end
}
