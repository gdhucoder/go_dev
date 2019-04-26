package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	// 模板全局变量
	myTemplate *template.Template
)

type Person struct {
	Name  string
	Title string
	age   string
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	p := Person{Name: "胡国栋", Title: "Engineer", age: "31"}
	myTemplate.Execute(w, p)
}

func initTemplate(filename string) (err error) {
	myTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	return
}

func main() {

	// 初始化模板
	initTemplate("D:\\project\\src\\go_dev\\day10\\template\\index.html")

	// 处理请求
	http.HandleFunc("/user/info", Info)

	//
	err := http.ListenAndServe("0.0.0.0:8890", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
