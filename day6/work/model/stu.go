package model

import "errors"

var (
	ErrBookNotFound = errors.New("book not found!")
)

// 学生信息管理功能，管理每个学生的姓名、年级、身份证、性别、
type Student struct {
	Name  string
	Grade int
	ID    string
	Sex   int
	books []*BorrowItem // 指针类型
}

type BorrowItem struct {
	book *Book
	num  int
}

func CreateStudent(name string, grade int, id string, sex int) (stu *Student) {
	stu = &Student{
		Name:  name,
		Grade: grade,
		ID:    id,
		Sex:   sex,
	}
	return
}

func (s *Student) AddBook(book *BorrowItem) {
	s.books = append(s.books, book)
}

func (s *Student) DelBook(b *BorrowItem) (err error) {
	for i := 0; i < len(s.books); i++ {
		if s.books[i].book.Name == b.book.Name {
			if s.books[i].num == b.num {
				// 删除slice
				front := s.books[0:i]
				left := s.books[i+1:]
				s.books = append(front, left...)
				return
			}
			s.books[i].num -= b.num
			return
		}
	}
	err = ErrBookNotFound
	return
}

func (s *Student) ListBook() []*BorrowItem {
	return s.books
}
