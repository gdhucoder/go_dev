package main

import "fmt"

type Car1 struct {
	Band  string
	Price float32
}

type Car2 struct {
	Band  string
	Price float32
}

type Train struct {
	Car1
	Car2
	Price float32
}

func main() {
	var t Train
	// t.Band = "sdf" // 这样使用会有歧义
	t.Car1.Band = "Tesla"
	t.Car1.Price = 7777.77
	t.Car2.Price = 123123
	t.Price = 9999.99
	fmt.Println(t)
	// {{Tesla 7777.77} { 123123} 9999.99}
}
