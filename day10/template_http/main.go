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

type Result struct {
	output string
}

func (p *Result) Write(b []byte) (n int, err error) {
	fmt.Println("called by template")
	p.output += string(b)
	return len(b), nil
}

type Person struct {
	Name  string
	Title string
	Age   int
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle info")
	p := Person{Name: "胡国栋", Title: "Engineer", Age: 17}
	myTemplate.Execute(w, p) // w实现了写的接口
	resultWriter := &Result{}
	myTemplate.Execute(resultWriter, p)
	fmt.Println(resultWriter.output) // 自己实现Write接口

	// called by template
	// called by template
	// called by template
	// called by template
	// called by template
	// called by template
	// called by template
	// <html>
	// 	<head>
	// 		<title>Engineer</title>
	// 	</head>
	// 	<body>
	// 		<h1>Hello, 胡国栋</h1>
	// 		{胡国栋 Engineer 31}
	// 	</body>
	// </html>
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
	initTemplate("D:\\project\\src\\go_dev\\day10\\template_http\\index.html")

	// 处理请求
	http.HandleFunc("/user/info", Info)

	//
	err := http.ListenAndServe("0.0.0.0:8890", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}
