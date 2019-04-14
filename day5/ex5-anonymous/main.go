package main

import (
	"fmt"
	"time"
)

type Car struct {
	Band  string
	Price float32
}

type Train struct {
	Car   // 匿名变量（字段）
	int   // 匿名变量
	Start time.Time
}

func main() {
	var t Train
	// t.Band = "Benz"
	t.Car.Band = "Audi"
	t.Price = 19999.99
	t.int = 50 // int 可以直接使用类型做为变量名称
	fmt.Println(t)
	// {{Benz 19999.99} 50 0001-01-01 00:00:00 +0000 UTC}
}
