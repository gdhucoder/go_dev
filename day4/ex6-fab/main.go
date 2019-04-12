package main

import "fmt"

// 使用非递归方式实现fab数列
// f(n) = f(n-1) + f(n-2)
// f2 = f1 + f0
// f3 = f2 + f1
func fab(n int) int {

	if n <= 1 {
		return 1
	}

	var a []int
	a = make([]int, n) // 新建立一个切片
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n-1] + a[n-2]

}

func main() {
	for i := 2; i < 1000; i++ {
		f := fab(i)
		fmt.Println(f)
	}

}
