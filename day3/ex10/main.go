package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanf("%s", &input) 
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("can not convert to int", err)
		return
	}
	fmt.Println(number)
}
