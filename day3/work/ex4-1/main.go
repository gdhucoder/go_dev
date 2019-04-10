package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			worldCount++
		case v >= '0' && v <= '9':
			numberCount++
		case v == ' ':
			spaceCount++
		default:
			otherCount++
		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("world count:%d\n space count:%d\n space count: %d\n other count: %d\n", wc, sc, nc, oc)

}
