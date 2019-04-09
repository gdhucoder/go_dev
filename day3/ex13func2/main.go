package main

import "fmt"

func calc(a, b int) (sum, avg int) {
	sum = a + b
	avg = sum / 2
	return
}

func main() {
	sum, _ := calc(10, 5)
	fmt.Println(sum)

}
