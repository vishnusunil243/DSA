package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
}
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	toRemove := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return toRemove
}
func (q *Queue) Print() {
	temp := Queue{}
	for !q.IsEmpty() {
		fmt.Printf("%d ", q.front.data)
		temp.Enqueue(q.Dequeue())
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}

func (s *Stack) Push(value int) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	toRemove := s.top.data
	s.top = s.top.next
	return toRemove
}
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.top.data
}
func (s *Stack) Print() {
	temp := Stack{}
	for !s.IsEmpty() {
		fmt.Printf("%d ", s.top.data)
		temp.Push(s.Pop())
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (s *Stack) insertAfter(value, target int) {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		if current == target {
			temp.Push(value)
		}
		temp.Push(current)

	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (q *Queue) DeleteWithValue(value int) {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		if current == value {
			continue
		}
		temp.Enqueue(current)
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (q *Queue) QueueToStack() *Stack {
	temp := Stack{}
	for !q.IsEmpty() {
		temp.Push(q.Dequeue())
	}
	return &temp
}
func (s *Stack) StackToQueue() *Queue {
	temp := Stack{}
	res := Queue{}
	for !s.IsEmpty() {
		temp.Push(s.Pop())
	}
	for !temp.IsEmpty() {
		res.Enqueue(temp.Pop())
	}
	return &res
}
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(12)
	s.Push(122)
	fmt.Println(s.Peek())
	s.insertAfter(22, 12)
	q := s.StackToQueue()
	q.Print()
	// q := Queue{}
	// q.Enqueue(1)
	// q.Enqueue(12)
	// q.Enqueue(222)
	// q.DeleteWithValue(1)
	// s := q.QueueToStack()
	// s.Print()
}
