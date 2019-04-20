package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	inputFile := "D:\\project\\src\\go_dev\\day7\\ex4-filebufio\\test.txt"
	outputFile := "D:\\project\\src\\go_dev\\day7\\ex4-filebufio\\test_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// File Error: open D:\project\src\go_dev\day7\ex4-filebufio\test1111.txt: The system cannot find the file specified.
		return
	}

	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0x644)
	if err != nil {
		panic(err.Error())
	}

}
