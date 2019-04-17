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

func testInt(a interface{}) {
	v := reflect.ValueOf(a)
	value := v.Int()
	fmt.Println(value)
}

// 通过反射修改值
func testModifyValue(a interface{}) {
	v := reflect.ValueOf(a) // 地址

	// the pointer v points to.
	v.Elem().SetInt(1000) // v.Elem() 相当于 (*ptr)地址取值操作
	fmt.Println(v.Elem().Int())

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

	testInt(123444)

	// 修改int类型的数据
	b := 100
	testModifyValue(&b)
	fmt.Println(b)
}
