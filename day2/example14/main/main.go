package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genRandInt(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(rand.Int31())
	}
}

// 小于100的10个随机整数
func genRanInt10() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(10))
	}
}

func genRanFloat() {
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Float64())
	}
}

// 初始化的
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	genRandInt(10)
	genRanInt10()
	genRanFloat()

	var ss = `adfadsf \n \n` // 原样输出
	fmt.Println(ss)

	//
	var string1 = `
	我是小熊熊，
	你是小naughty，
	你要不听话，
	我也没办法。
	`
	fmt.Println(string1)
}
