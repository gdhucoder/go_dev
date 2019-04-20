package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile := "D:\\project\\src\\go_dev\\day7\\ex4-filebufio\\test.txt"
	dstFile := "D:\\project\\src\\go_dev\\day7\\ex7-copy\\test_copy.txt"
	CopyFile(dstFile, inputFile)
	fmt.Println("Copy Success!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0664)

	if err != nil {
		return
	}

	defer dst.Close()
	return io.Copy(dst, src)
}
