package main

import (
	"fmt"
	"go_dev/day1/package_example/calc"
)

func main() {
	a := calc.Add(10, 2)
	fmt.Println("a=", a)
	b := calc.Sub(20, 19)
	fmt.Println("b=", b)

}
