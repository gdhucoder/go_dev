package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
}

// 从字符串里面格式化输入
// 适合处理一行文本数据
func main() {
	var stu Student
	str := "xiaoming 18 98.5"
	fmt.Sscanf(str, "%s %d %f", &stu.Name, &stu.Age, &stu.Score)
	fmt.Println(stu)
	// {xiaoming 18 98.5}
}
