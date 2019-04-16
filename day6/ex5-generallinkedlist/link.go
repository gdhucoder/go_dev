package main

import "fmt"

// 写一个通用链表

// 链表结点
type LinkNode struct {
	data interface{} // 任何数据类型
	next *LinkNode
}

// 链表
type Link struct {
	head *LinkNode
	tail *LinkNode
}

type Student struct {
	Name string
}

func (l *Link) InsertHead(data interface{}) {
	var node = &LinkNode{
		data: data,
	}
	if l.head == nil && l.tail == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head = node
}

func (l *Link) InsertTail(data interface{}) {
	var node = &LinkNode{
		data: data,
	}
	if l.head == nil && l.tail == nil {
		l.head = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *Link) Traverser() {
	p := l.head
	for p != nil {
		fmt.Println(p.data)
		p = p.next
	}
}
