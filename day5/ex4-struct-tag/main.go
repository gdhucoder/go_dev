package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"student_name"`
	Age   int    `json:"student_age"`
	Score int    `json:"student_score"`
}

func main() {
	var stu Student = Student{
		Name:  "huguodong",
		Age:   33,
		Score: 80,
	}

	data, err := json.Marshal(stu)
	// 由于json是外部的包，访问不到这个包中struct中小写字母的
	// 变量。有两种方法，第一是改结构体中的变量为大写。
	// 第二是使用tag

	if err != nil {
		fmt.Println("json encode stu failed, err: ", err)
		return
	}

	fmt.Println(string(data))
}
