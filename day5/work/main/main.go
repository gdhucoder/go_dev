package main

import (
	"fmt"
	"go_dev/day5/work/model" // 自定义的包
)

func main() {
	var allBooks []model.Book
	// allBooks = make([]book.Book, 5)
	// name string, author string, issuedate string, copy int
	var book = model.CreateNewBook("alg", "SZ", "2019-04-15", 5)
	fmt.Println(*book)

	fmt.Println(len(allBooks))
	model.InitBook()
}
