package main

func main() {
	var link Link
	for i := 0; i < 10; i++ {
		// link.InsertTail(i)
		link.InsertHead(i)
	}
	// link.InsertTail("huguodong")
	link.InsertHead("huguodong")
	link.Traverser()
	// var stu Student
	// fmt.Println(stu)
}
