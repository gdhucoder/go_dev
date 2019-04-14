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

// 不能直接传入head的副本
// 想修改head的值需要传入head的指针
// 二级指针就够复杂了
func insertHead_V2(p **Student) {
	for i := 0; i < 10; i++ {
		var stu *Student = &Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = stu
		fmt.Printf(" head address %p\n", *p)
	}
}

// 可能涉及到头结点变动的问题
func del(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prev.next = p.next
			break
		} else {
			prev = p
			p = p.next
		}
	}
}

// 可能涉及到头结点变动的问题
func insert(p *Student) {
	for p != nil {
		if p.Name == "stu5" {
			stu := &Student{
				Name:  "huguoodng",
				Age:   33,
				Score: 999,
			}
			stu.next = p.next
			p.next = stu
			break
		} else {
			p = p.next
		}
	}
}

func main() {
	var head *Student = &Student{
		Name:  "wtw",
		Age:   1,
		Score: 100,
	}

	// tail := head

	// insertTail(tail)

	// head = insertHead(head)

	insertHead_V2(&head)

	traverse(head)

	del(head)

	traverse(head)

	insert(head)

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
