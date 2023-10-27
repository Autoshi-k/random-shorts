package main

import "fmt"

func testDataStructures() {
	// Implement a stack or a queue using slices in Go.
	stack := Stack{}
	fmt.Printf("Created new stack, len: %d\n", stack.Len())

	stack.Push(Item{Id: "1"})
	fmt.Printf("Added item 1, len: %d\n", stack.Len())
	stack.Push(Item{Id: "2"})
	fmt.Printf("Added item 2, len: %d\n", stack.Len())
	stack.Push(Item{Id: "3"})
	fmt.Printf("Added item 3, len: %d\n", stack.Len())

	fmt.Printf("Popped item, id: %s, len: %d\n", stack.Pop().Id, stack.Len())
	fmt.Printf("Popped item, id: %s, len: %d\n", stack.Pop().Id, stack.Len())
	fmt.Printf("Popped item, id: %s, len: %d\n", stack.Pop().Id, stack.Len())
	fmt.Printf("Popped item, id: %s, len: %d\n", stack.Pop().Id, stack.Len())

	stack.Push(Item{Id: "4"})
	fmt.Printf("Added item 4, len: %d\n", stack.Len())
	stack.Push(Item{Id: "5"})
	fmt.Printf("Added item 5, len: %d\n", stack.Len())
	fmt.Printf("Popped item, id: %s, len: %d\n", stack.Pop().Id, stack.Len())
	stack.Push(Item{Id: "6"})
	fmt.Printf("Added item 6, len: %d\n", stack.Len())

	fmt.Println("")
	queue := Queue{}
	fmt.Printf("Created new queue, len: %d\n", queue.Len())

	queue.Enqueue(Item{Id: "1"})
	fmt.Printf("Added item 1, len: %d\n", queue.Len())
	queue.Enqueue(Item{Id: "2"})
	fmt.Printf("Added item 2, len: %d\n", queue.Len())
	queue.Enqueue(Item{Id: "3"})
	fmt.Printf("Added item 3, len: %d\n", queue.Len())

	fmt.Printf("Dequeue item, id: %s, len: %d\n", queue.Dequeue().Id, queue.Len())
	fmt.Printf("Dequeue item, id: %s, len: %d\n", queue.Dequeue().Id, queue.Len())
	fmt.Printf("Dequeue item, id: %s, len: %d\n", queue.Dequeue().Id, queue.Len())
	fmt.Printf("Dequeue item, id: %s, len: %d\n", queue.Dequeue().Id, queue.Len())

	queue.Enqueue(Item{Id: "4"})
	fmt.Printf("Added item 4, len: %d\n", queue.Len())
	queue.Enqueue(Item{Id: "5"})
	fmt.Printf("Added item 5, len: %d\n", queue.Len())
	fmt.Printf("Dequeue item, id: %s, len: %d\n", queue.Dequeue().Id, queue.Len())
	queue.Enqueue(Item{Id: "6"})
	fmt.Printf("Added item 6, len: %d\n", queue.Len())

	// Create a custom data structure, like a linked list or a hash table.
	// Implement a binary search tree and perform basic operations like insertion, deletion, and searching.
}

type Item struct {
	Id string
}

type Stack []Item

func (s *Stack) Push(item Item) {
	*s = append(*s, item)
}

func (s *Stack) Pop() Item {
	stack := *s

	if len(stack) == 0 {
		return Item{}
	}

	res := stack[len(stack)-1]
	*s = stack[:len(stack)-1]

	return res
}

func (s *Stack) Len() int {
	return len(*s)
}

type Queue []Item

func (q *Queue) Enqueue(item Item) {
	*q = append(*q, item)
}

func (q *Queue) Dequeue() Item {
	queue := *q

	if len(queue) == 0 {
		return Item{}
	}

	res := queue[0]
	*q = queue[1:]

	return res
}

func (q *Queue) Len() int {
	return len(*q)
}

//type iStack interface {
//	Push(item Item)
//	Pop() Item
//	Len() int
//}
//
//type Stack struct {
//	current *StackNode
//	count   int
//}
//
//type StackNode struct {
//	value    Item
//	previous *StackNode
//}
//
//func (s *Stack) Push(item Item) {
//	s.current = &StackNode{value: item, previous: s.current}
//	s.count++
//}
//
//func (s *Stack) Pop() Item {
//	if s.count == 0 {
//		return Item{}
//	}
//
//	removedNode := s.current.value
//
//	s.current = s.current.previous
//	s.count--
//
//	return removedNode
//}
//
//func (s *Stack) Len() int {
//	return s.count
//}
//
//type iQueue interface {
//	Enqueue(item Item)
//	Dequeue() Item
//	Len() int
//}
//type Queue struct {
//	endQ   *QueueNode // take from end
//	startQ *QueueNode // insert to start
//	count  int
//}
//
//type QueueNode struct {
//	value Item
//	front *QueueNode
//	back  *QueueNode
//}
//
//func (q *Queue) Enqueue(item Item) {
//	qNode := &QueueNode{value: item, front: q.startQ}
//	if q.startQ != nil {
//		q.startQ.back = qNode
//	}
//
//	q.startQ = qNode
//
//	if q.endQ == nil {
//		q.endQ = q.startQ
//	}
//
//	q.count++
//}
//
//func (q *Queue) Dequeue() Item {
//	if q.count == 0 {
//		return Item{}
//	}
//
//	res := q.endQ.value
//	q.endQ = q.endQ.back
//
//	if q.count == 1 {
//		q.startQ = nil
//	}
//
//	q.count--
//	return res
//}
//
//func (q *Queue) Len() int {
//	return q.count
//}
