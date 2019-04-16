package model

type Student struct {
	Id            string // 学号
	Name          string //
	Grade         int
	IDNum         string       // 身份证号
	Sex           int          // 0 female 1 male
	BooksBorrowed []BookRecord // 借了几本书
}

type BookRecord struct {
	book *Book
	num  int
}

func InitStudent() []Student {
	return nil
}

func CreateNewStudent(name string) *Student {
	var stu *Student = &Student{
		Name: name,
	}
	return stu
}

// 查询书籍
func (stu *Student) QueryBooks(bookName string) []Book {
	return nil
}

// 借阅书籍 num几本
func (stu *Student) BorrowBook(book *Book, num int) {

}

func (stu *Student) ReturnBook(book *Book, num int) {

}
