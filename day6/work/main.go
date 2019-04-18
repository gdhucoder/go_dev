package main

import (
	"fmt"
	"go_dev/day6/work/model"
)

var (
	books    []*model.Book
	students []*model.Student
	users    []*model.User
)

func init() {
	books = InitBooks()
	users, students = InitUsers()
}

func InitStudent() (students []*model.Student) {
	students = make([]*model.Student, 2)
	students[0] = model.CreateStudent("wtw", 1, "220221198909260202", 0)
	students[1] = model.CreateStudent("hgd", 2, "220221198909260203", 1)
	return
}

func InitUsers() (users []*model.User, students []*model.Student) {
	students = make([]*model.Student, 2)
	users = make([]*model.User, 2)
	users[0] = model.Register("hgd", "123", "hgd", 2, "220221198909260203", 1)
	users[1] = model.Register("wtw", "123", "wtw", 1, "220221198909260202", 0)
	students[0] = users[0].StudentInfo
	students[1] = users[1].StudentInfo
	return
}

// 初始化书库
func InitBooks() (books []*model.Book) {
	books = make([]*model.Book, 2)
	books[0] = model.CreateBook("人类简史", 3, "赫拉利", "2019-04-18")
	books[1] = model.CreateBook("硅谷之谜", 10, "吴军", "2019-02-13")
	books = append(books, model.CreateBook("硅谷之谜1", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜2", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜3", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜4", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜5", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜6", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜7", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜8", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜9", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜10", 10, "吴军", "2019-02-13"))
	books = append(books, model.CreateBook("硅谷之谜11", 10, "吴军", "2019-02-13"))
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

	model.Top10Books(books)

	err = model.Login(users, "hgd", "123")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user login success")
}
