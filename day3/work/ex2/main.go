package main

import "fmt"

// 一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
// 编程找出1000以内的所有完数。
func isPerfectNum(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum == n {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 2; i < 1000; i++ {
		if isPerfectNum(i) {
			fmt.Println(i)
		}
	}
}
