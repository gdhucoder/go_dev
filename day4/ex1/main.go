package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error) {
	return errors.New("init config error")
}

func test() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// b := 0
	// a := 1 / b
	// fmt.Println(a)
	err := initConfig()

	if err != nil {
		panic(err)
	}

}

func main() {

	test()
	var i int
	fmt.Println(i)
	j := new(int)
	*j = 100
	fmt.Println(j)
	fmt.Println(*j)
	// new 出来的分配内存地址
	// 0
	// 0xc000058088

	var a []int // 切片
	a = append(a, 10, 11, 23)
	a = append(a, a...)
	fmt.Println(a)
	// [10 11 23 10 11 23]
}
