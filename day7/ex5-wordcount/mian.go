package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type WordCount struct {
	Letter     int
	Num        int
	WhiteSpace int
	Other      int
}

func main() {

	var cnt WordCount

	file, err := os.Open("D:\\project\\src\\go_dev\\day7\\ex4-filebufio\\test.txt")
	if err != nil {
		fmt.Printf("open file error : %v\n", err)
	}

	reader := bufio.NewReader(file)

	for {
		// 读取一行
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		strArr := []rune(str)
		for _, v := range strArr {
			switch {
			case v >= 'A' && v <= 'Z':
				fallthrough
			case v >= 'a' && v <= 'z':
				cnt.Letter++
			case v >= '0' && v <= '9':
				cnt.Num++
			case v == ' ' || v == '\t':
				cnt.WhiteSpace++
			default:
				cnt.Other++
			}
		}
	}
	fmt.Printf("Num of letter: %d\n", cnt.Letter)
	fmt.Printf("Num of number: %d\n", cnt.Num)
	fmt.Printf("Num of whitespace: %d\n", cnt.WhiteSpace)
	fmt.Printf("Num of other: %d\n", cnt.Other)
	// Num of letter: 522
	// Num of number: 4
	// Num of whitespace: 193
	// Num of other: 318
}
