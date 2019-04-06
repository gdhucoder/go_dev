package main

import "time"

/*
容易找 使用main.go 文件
*/
func main() {
	for i := 0; i < 100; i++ {
		go test_goroutine(i)
	}
	time.Sleep(time.Second)
}
