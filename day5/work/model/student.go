package student

import (
	"go_dev/day5/work/book"
)

type Student struct {
	Id            string // 学号
	Name          string //
	Grade         int
	IDNum         string      // 身份证号
	Sex           int         // 0 female 1 male
	BooksBorrowed []book.Book // 借了几本书
}

func InitStudent() []Student {
	var students []Student

}
