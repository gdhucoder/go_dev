package model

import "errors"

var (
	ErrStockNotEnough = errors.New("Stock is not enough!")
)

// 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期
type Book struct {
	Name       string
	Total      int
	Author     string
	CreateDate string
}

func CreateBook(name string, total int, author, createdate string) (book *Book) {
	book = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateDate: createdate,
	}
	return
}

func (b *Book) canBorrow(c int) bool {
	return b.Total >= c
}

func (b *Book) Borrow(c int) (err error) {
	if !b.canBorrow(c) {
		err = ErrStockNotEnough
	}
	b.Total -= c
	return
}

func (b *Book) ReturnBook(c int) {
	b.Total += c
}
