package main

import "fmt"

func main() {
	var mapslice = make([]map[int]string, 2)
	mapslice[0] = make(map[int]string)
	mapslice[0][1] = "123"
	fmt.Println(mapslice)
}
