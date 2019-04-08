package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// * This is known as "dereferencing" or "indirecting".
// * 表示解引用
func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
