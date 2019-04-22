package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	var intChan chan int
	intChan = make(chan int, 1)
	intChan <- 10

	b := <-intChan
	fmt.Println(b)

	stu := &Student{
		Name: "wtw",
	}

	// 存放任意数据类型
	var stuChan chan interface{}
	stuChan = make(chan interface{}, 1)
	stuChan <- stu

	stu2 := <-stuChan
	fmt.Println(stu2)

	var stu03 *Student
	stu03, ok := stu2.(*Student) // interface 转换成Student
	if !ok {
		fmt.Println("can not convert")
	}
	fmt.Println("stu03:", *stu03)
}
