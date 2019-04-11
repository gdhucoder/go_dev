package main

import "fmt"

func add() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

/* 这个两个很相似
class A{
	int x
	public int add(int d){
		x = x+d
		return x
	}
}

*/

// 闭包：一个函数和与其相关的引用环境组合而成的实体
// a closure is a record storing a function together with an environment
// https://gobyexample.com/closures

func main() {
	// 闭包
	f := add() // f 是一个函数了

	fmt.Println(f(1)) // x 的值会被保存下来
	fmt.Println(f(100))
	fmt.Println(f(1000))
}
