package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("D:\\project\\src\\go_dev\\day7\\ex2-fileio\\test.log",
		os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("error :", err)
	}

	fmt.Fprintf(file, "this is a log message.\n")
	fmt.Fprintf(file, "今天外面下的雨好大！\n")
	defer file.Close()
}
