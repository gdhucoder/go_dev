package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name  string
	Title string
	age   string
}

func main() {
	// src\go_dev\day10\template\index.html
	t, err := template.ParseFiles("D:\\project\\src\\go_dev\\day10\\template\\index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "胡国栋", Title: "Engineer", age: "31"}
	// b标准输出可以改为文件或者response
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
