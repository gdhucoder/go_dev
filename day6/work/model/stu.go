package model

import (
	"fmt"
	"time"
)

var (
// ErrBookNotFound = errors.New("book not found!")
)

const (
	MAXDAYS = 5 // 最大期限30天
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
	book               *Book
	num                int
	startBorrowDate    string // 借书的日期
	expectedReturnDate string // 预计还书日期
	isExpired          bool   // 是否已经过期
	borrowedDays       int
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

// 学生借书
func (s *Student) BorrowBook(book *Book, num int) (err error) {
	err = book.Borrow(s, num)
	if err != nil {
		return
	}
	start, end := StartEndDate()
	borrowItem := &BorrowItem{
		book:               book,
		num:                num,
		startBorrowDate:    start,
		expectedReturnDate: end,
	}

	s.books = append(s.books, borrowItem)
	return
}

func StartEndDate() (start, end string) {
	start = time.Now().Format(DateFormat)
	end = time.Now().AddDate(0, 0, MAXDAYS).Format(DateFormat)
	return
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

func (s *Student) PrintBorrowedBooks() {
	for i, v := range s.books {
		fmt.Printf("%d %v %d\n", i, *(v.book), v.num)
	}
}

// 更新借书同学借书状态
func RollDate(students []*Student) {
	for _, stu := range students {
		if len(stu.books) > 0 { // 已经借书
			for _, item := range stu.books {
				item.borrowedDays++
				if item.borrowedDays > MAXDAYS { // expired
					item.isExpired = true // 过期
					fmt.Printf("student name: %s, borrowed book: %s has expired! expected return date: %s, has expired %d days\n",
						stu.Name, item.book.Name, item.expectedReturnDate, (item.borrowedDays - MAXDAYS))

				}
			}
		}
	}
}
