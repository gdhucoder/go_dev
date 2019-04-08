package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	// The special prefix & returns a pointer to the struct value.
	p = &Vertex{2, 4}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
