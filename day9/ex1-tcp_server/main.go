package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf) // 读取n个
		if err != nil {
			fmt.Println("Read err: ", err)
			return
		}
		fmt.Println("read: ", string(buf[0:n]))
	}
}

func main() {

	fmt.Println("Start Server...")

	// 监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen failed, err: ", err)
		return
	}

	for {

		// 建立与客户端的连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err: ", err)
			continue
		}

		// 建立goroutine处理该连接
		go process(conn)
	}

}
