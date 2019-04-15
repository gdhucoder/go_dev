package main

import (
	"fmt"
	"go_dev/day5/work/book" // 自定义的包
)

func main() {
	var allBooks []book.Book
	// allBooks = make([]book.Book, 5)

	fmt.Println(len(allBooks))
	book.InitBook()
}
