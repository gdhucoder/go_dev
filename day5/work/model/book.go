package book

import "fmt"

type Book struct {
	ISBN      string // isbn编号
	Name      string // 书名
	Copy      int    // 副本
	OutNum    int    // 出借几本
	InNum     int    // 在库几本
	Author    string // 作者
	IssueDate string // 出版日期
}

func (b *Book) GetName() string {
	return b.Name
}

// 出借
func (b *Book) LendBook() {
	b.InNum++
	b.OutNum--
}

// 还书
func (b *Book) ReturnBook() {
	b.InNum--
	b.OutNum++
}

// 新增副本
func (b *Book) AddNewCopy() {
	b.InNum++
	b.Copy++
}

func (b Book) String() string {
	str := fmt.Sprintf("<ISBN: %s,\n BOOKNAME: %s,\n COPY: %d, IN %d, OUT %d,\nISSUEDATE: %s>\n",
		b.ISBN, b.Name, b.Copy, b.InNum, b.OutNum, b.IssueDate)
	return str
}

// x

/*

初始化书库
使用slice维护书
books = []
append(books, newbook) // 新增书

1 AddNewBook(book) return []
2 FindBookByName(bookeName) return []
3 FindBookByAuthor(author) return []
4 FindBookByIssueDate(start, end) return []

*/

// 初始化书库
func InitBook() []Book {
	var books []Book
	for i := 0; i < 10; i++ {
		var nbk Book = Book{
			ISBN:      "192-22",
			Name:      "算法4",
			Author:    "RS",
			IssueDate: "2016-01-01",
			Copy:      1,
			InNum:     1,
		}
		books = AddNewBook(books, nbk) // 新增一本书
	}

	for i := 0; i < 10; i++ {
		var nbk Book = Book{
			ISBN:      "192-23",
			Name:      "人类简史",
			Author:    "赫拉利",
			IssueDate: "2012-02-01",
			Copy:      1,
			InNum:     1,
		}
		books = AddNewBook(books, nbk) // 新增一本书
	}

	fmt.Println(books)
	return books
}

func AddNewBook(books []Book, newBook Book) []Book {
	res := FindBookByName(books, newBook.Name)
	if len(res) == 0 { // 不存在
		res = append(books, newBook)
		return res
	} else { // 存在
		res[0].AddNewCopy()
		return books
	}

}

func FindBookByName(books []Book, name string) []Book {
	var res []Book
	for i, v := range books {
		if v.Name == name {
			res = books[i : i+1]
			break
		}
	}
	return res
}

func FindBookByAuthor(books []Book, author string) []Book {
	return nil
}

func FindBookByIssueDate(books []Book, start string, end string) []Book {
	return nil
}
