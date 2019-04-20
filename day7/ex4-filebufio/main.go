package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 打开文件
	file, err := os.Open("D:\\project\\src\\go_dev\\day7\\ex4-filebufio\\test.txt")
	if err != nil {
		fmt.Printf("open file error: %s\n", err)
	}

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	// 带缓冲区的文件读写
}
