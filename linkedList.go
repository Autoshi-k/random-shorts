package main

import "fmt"

type LinkedList struct {
	head *Node
}

type Node struct {
	data interface{}
	next *Node
}

func (l *LinkedList) Append(value interface{}) {
	node := &Node{
		data: value,
	}

	n := l.head

	if n == nil {
		l.head = node
		return
	}

	for n.next != nil {
		n = n.next
	}

	n.next = node
}

func (l *LinkedList) Remove(value interface{}) {
	n := l.head

	if n.data == value {
		l.head = l.head.next
		return
	}

	for n.next != nil {
		if n.next.data == value {
			n.next = n.next.next
			return
		}
		n = n.next
	}
}

func (l *LinkedList) Print() {

	n := l.head
	var prnt string
	var index int

	for n != nil {
		index++
		prnt += fmt.Sprintf(" [%d] %+v\n", index, n.data)
		n = n.next
	}

	fmt.Println(prnt)
}
