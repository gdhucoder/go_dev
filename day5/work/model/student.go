package model

type Student struct {
	Id            string // 学号
	Name          string //
	Grade         int
	IDNum         string // 身份证号
	Sex           int    // 0 female 1 male
	BooksBorrowed []Book // 借了几本书
}

func InitStudent() []Student {
	var students []Student

}
