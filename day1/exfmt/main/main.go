package main

import "fmt"

func printBinary(num int) {
	fmt.Printf("二进制%b\n", num)
}

func printDecimal(num int) {
	fmt.Printf("十进制%d\n", num)
}

func printHexical(num int) {
	fmt.Printf("16进制小写%x\n", num)
	fmt.Printf("16进制大写%X\n", num)
}

func printOct(num int) {
	fmt.Printf("八进制%o\n", num)
}

func printFloat(num float32) {
	// 保留两位小数
	fmt.Printf("%.2f\n", num)
}

func main() {
	num := 15
	printBinary(num)
	printDecimal(num)
	printOct(num)
	printHexical(num)
	printFloat(1.66666)
}
