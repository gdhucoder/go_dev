package main

import (
	"fmt"
	"go_dev/day6/work/model"
)

var (
	books    []*model.Book
	students []*model.Student
)

func init() {
	books = InitBooks()
	students = InitStudent()
}

func InitStudent() (students []*model.Student) {
	students = make([]*model.Student, 2)
	students[0] = model.CreateStudent("wtw", 1, "220221198909260202", 0)
	students[1] = model.CreateStudent("hgd", 2, "220221198909260203", 1)
	return
}

func InitBooks() (books []*model.Book) {
	books = make([]*model.Book, 2)
	books[0] = model.CreateBook("人类简史", 3, "赫拉利", "2019-04-18")
	books[1] = model.CreateBook("硅谷之谜", 10, "吴军", "2019-02-13")
	return
}

// 初始化用户
func InitUsers() {
	return
}

func main() {
	fdBook, err := model.FindBookByName(books, "人类简史")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Printf("book found %s\n", fdBook)
	err = students[0].BorrowBook(fdBook[0], 2)
	err = students[1].BorrowBook(fdBook[0], 1)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	// booksBorrowed := students[1].ListBook()
	students[1].PrintBorrowedBooks()
	fdBook[0].BorrowedByStudents()

}
