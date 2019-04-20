package main

import (
	"bufio"
	"fmt"
	"os"
)

// 带缓存区的读
func main() {

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("读入数据:%s\n", str)
	// hello world!
	// 读入数据:hello world!
}
