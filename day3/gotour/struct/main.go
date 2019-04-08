package main

import "fmt"

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{100, 1101010})
}
