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

	fmt.Println("")

	list := &LinkedList{}

	list.Append(1)
	list.Print()
	list.Append(2)
	list.Print()
	list.Append(4)
	list.Print()
	list.Remove(4)
	list.Print()
	list.Append(3)
	list.Print()
	list.Append(5)
	list.Print()
	list.Remove(1)
	list.Print()
	list.Append(4)
	list.Print()
	list.Remove(3)
	list.Print()

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
