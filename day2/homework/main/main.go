package main

import (
	"fmt"
	"math"
)

// 101 到 200 之间的质数
func test1() {
	for i := 101; i <= 200; i++ {
		if isPrime(i) {
			fmt.Printf("%d is Prime.\n", i)
		}
	}
}

// 打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
// 立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
// 方＋5 的三次方＋3 的三次方。

func test2() {
	for i := 100; i < 1000; i++ {
		if isShuiXian(i) {
			fmt.Println(i)
		}
	}
}

// 3. 对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
// sum(n) = sum(n-1) + f(n)
func test3(n int) {
	res := sum(n)
	fmt.Printf("res %d\n", res)
}

//  对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
func sum(n int) int {

	// n = 1
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1 + 2
	}

	// n = 2
	// sum(2) = 1! + (1!*2)

	// n = 3
	// sum(3) = sum(2) + (sum(2) - sum(1))*3
	return sum(n-1) + (sum(n-1)-sum(n-2))*n
}

func isPrime(n int) bool {
	upper := int(math.Sqrt(float64(n)))
	for i := 2; i <= upper; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isShuiXian(n int) bool {
	first := n / 100      // 第一位
	second := n / 10 % 10 // 第二位
	third := n % 10
	// fmt.Println(first, second, third)
	if (first*first*first + second*second*second + third*third*third) == n {
		return true
	} else {
		return false
	}
}

func main() {
	isShuiXian(153)

	test3(4)

	res := isPrime(4)
	fmt.Println(res)
	test1()
	test2()

}
