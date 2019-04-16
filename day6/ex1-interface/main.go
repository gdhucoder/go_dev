package main

import "fmt"

type Test interface {
}

func main() {
	var a interface{} // 空接口
	var b int
	a = b
	fmt.Printf("type of a is %T \n", a) //type of a is int
}
