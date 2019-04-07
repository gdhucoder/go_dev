package main

import (
	"fmt"
	"math/rand"
)

func genRandInt(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(rand.Int31())
	}
}

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

func main() {
	rand.Seed(0)
	genRandInt(10)
	genRanInt10()
	genRanFloat()
}
