package main

import "fmt"

type Test interface {
}

type CarI interface {
	Run()
	Name() string
	DiDi()
}

func main() {
	var a interface{} // 空接口
	var b int
	a = b                               // int 类型
	fmt.Printf("type of a is %T \n", a) //type of a is int

	var c float32
	a = c
	fmt.Println(c) //float32
}
