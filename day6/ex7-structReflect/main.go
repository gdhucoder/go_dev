package main

import (
	"fmt"
	"reflect"
)

// Student
type Student struct {
	Name  string
	Age   int
	Score float64
}

func (s *Student) Print() {
	fmt.Println("---------start-------------")
	fmt.Println(s)
	fmt.Println("---------end-------------")
}

func (s *Student) Set(name string, age int, socre float64) {
	s.Name = name
	s.Age = age
	s.Score = socre
	fmt.Println(s)
}

func testStuReflect(a interface{}) {
	v := reflect.ValueOf(a)
	kd := v.Elem().Kind()
	fmt.Println(kd)
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	numOfFields := v.Elem().NumField() // v 是地址，通过地址获得 *ptr，获取结构体中的字段名
	fmt.Printf("struct has %d fields\n", numOfFields)
	numOfMethods := v.NumMethod() // 获取 结构体中的指针方法
	fmt.Printf("struct has %d methods\n", numOfMethods)

	for i := 0; i < numOfFields; i++ {
		fmt.Println(v.Elem().Field(i))
		// huguodong
		// 33
		// 99.9
	}

	// 设置参数 v.Method(1).Call(params) params 是切片[]reflect.Value 类型
	params := make([]reflect.Value, 3) // 申明3个大小的切片
	params[0] = reflect.ValueOf("wtw") // 通过反射变成Value类型
	params[1] = reflect.ValueOf(1)
	params[2] = reflect.ValueOf(988.8)

	v.Method(1).Call(params)          // 通过方法的序号调用方法
	v.MethodByName("Print").Call(nil) // 通过名字调用方法
}

func main() {
	var stu = &Student{
		Name:  "huguodong",
		Age:   33,
		Score: 99.9,
	}

	// 通过反射修改值
	testStuReflect(stu)
	fmt.Println(stu)
}
