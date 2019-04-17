package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func testReflect(p interface{}) {
	t := reflect.TypeOf(p) // type 指的是具体的类型
	v := reflect.ValueOf(p)
	fmt.Println(t, v)
	k := v.Kind() // 指的是类别
	fmt.Println(k)
	vi := v.Interface()     // 转成interface
	stu, ok := vi.(Student) // 转成student
	if ok {
		fmt.Println(stu)
		// {huguodong 33 99.9}
	}
}

func main() {
	// a := 2
	// testReflect(a)

	var stu = Student{
		Name:  "huguodong",
		Age:   33,
		Score: 99.9,
	}
	testReflect(stu)

}
