package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// 建立连接
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error dialing ", err.Error())
	}
	defer conn.Close() // 关闭连接

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimedInput := strings.Trim(input, "\r\n")
		if trimedInput == "Q" {
			return
		}

		// 数据发送
		_, err = conn.Write([]byte(trimedInput))
		if err != nil {
			return
		}
	}
}
