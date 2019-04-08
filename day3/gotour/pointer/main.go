package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i
	fmt.Println(*p) // read i through pointer

	*p = 21
	fmt.Println(i) // set i through the pointer

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j) // see the new value of j
}
