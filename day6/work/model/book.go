package model

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrStockNotEnough = errors.New("Stock is not enough!")
	ErrBookNotFound   = errors.New("Book is not found!")
	ErrRecordNotFound = errors.New("Borrow Record is not found!")
	ErrReturnNumWrong = errors.New("Return Num is Wrong!")
)

// 书籍录入功能，书籍信息包括书名、副本数、作者、出版日期
type Book struct {
	Name          string
	Total         int // 一共有几本
	Author        string
	CreateDate    string
	InStock       int // 库存
	Out           int // 借出几本
	Borrowed      []*Student
	BorrowedTimes int // 书籍被借过多少次
}

type ByBorrowedTimes []*Book // 按照被借的次数排序

// 实现排序接口
func (b ByBorrowedTimes) Len() int {
	return len(b)
}

func (b ByBorrowedTimes) Less(i, j int) bool {
	return b[i].BorrowedTimes < b[j].BorrowedTimes
}

func (b ByBorrowedTimes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// 书籍录入功能
// 返回一个书的指针
func CreateBook(name string, total int, author, createdate string) (book *Book) {
	book = &Book{
		Name:       name,
		Total:      total,
		Author:     author,
		CreateDate: createdate,
		InStock:    total,
		Out:        0, // 初始默认是0
	}
	return
}

func (b *Book) canBorrow(c int) bool {
	return b.InStock >= c
}

// 借书
func (b *Book) Borrow(s *Student, c int) (err error) {
	if !b.canBorrow(c) {
		err = ErrStockNotEnough
	}
	b.InStock -= c // 库存还剩几本
	b.Out += c     // 在外几本
	b.Borrowed = append(b.Borrowed, s)
	b.BorrowedTimes++
	return
}

// 还书
func (b *Book) ReturnBook(s *Student, c int) (err error) {

	if b.Out <= c {
		err = ErrReturnNumWrong
		return
	}

	b.Out -= c
	b.InStock += c
	for i, v := range b.Borrowed {
		// 删除学生列表中的数据
		if v.Name == s.Name {
			pre := b.Borrowed[0:i]
			left := b.Borrowed[i+1:]
			b.Borrowed = append(pre, left...)
			return
		}
	}
	err = ErrRecordNotFound
	return
}

// 通过书名查找书
func FindBookByName(books []*Book, name string) (res []*Book, err error) {
	for i, v := range books {
		if v.Name == name {
			res = books[i : i+1]
			return
		}
	}
	err = ErrBookNotFound
	return
}

func (b *Book) BorrowedByStudents() []*Student {
	for i, v := range b.Borrowed {
		fmt.Println(b.Name, b.InStock, i, *v)
	}
	return b.Borrowed
}

// 增加显示热门图书的功能，被借次数最多的top10
func Top10Books(books []*Book) (res []*Book) {
	// 将书籍按照被借的次数倒序排列
	var btimes ByBorrowedTimes
	btimes = books
	sort.Sort(sort.Reverse(btimes))
	for i, b := range btimes {
		fmt.Println(i, b.Name, b.BorrowedTimes)
	}
	numBooks := len(btimes)
	if numBooks > 10 {
		res = btimes[0:10]
	} else {
		res = btimes[0:numBooks]
	}
	// fmt.Println(btimes)
	return
}

//
