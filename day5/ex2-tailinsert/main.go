package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func traverse(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func insertTail(p *Student) {
	tail := p

	for i := 0; i < 10; i++ {
		var stu *Student = &Student{
			Name:  fmt.Sprintf("%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}
}

// p 指针变量的副本
func insertHead(p *Student) *Student {
	head := p
	for i := 0; i < 10; i++ {
		var stu *Student = &Student{
			Name:  fmt.Sprintf("%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = head
		head = stu
		fmt.Printf(" head address %p\n", head)
	}
	return head
}

func main() {
	var head *Student = &Student{
		Name:  "wtw",
		Age:   1,
		Score: 100,
	}

	// tail := head

	// insertTail(tail)

	head = insertHead(head)

	traverse(head)

	// {wtw 1 100 0xc00006a300}
	// {0 81 94.05091 0xc00006a330}
	// {1 47 43.77142 0xc00006a360}
	// {2 81 68.682304 0xc00006a390}
	// {3 25 15.651925 0xc00006a3c0}
	// {4 56 30.091187 0xc00006a3f0}
	// {5 94 81.36399 0xc00006a420}
	// {6 62 38.06572 0xc00006a450}
	// {7 28 46.888985 0xc00006a480}
	// {8 11 29.310184 0xc00006a4b0}
	// {9 37 21.855305 <nil>}

}
