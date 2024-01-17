package main

import "fmt"

type Node struct {
	data int
	next *Node
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
	tmp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		fmt.Printf("%d ", current)
		tmp.Enqueue(current)
	}
	for !tmp.IsEmpty() {
		q.Enqueue(tmp.Dequeue())
	}
}
func NewStack() *Stack {
	return &Stack{
		q1: &Queue{},
		q2: &Queue{},
	}
}

type Stack struct {
	q1, q2 *Queue
}

func (s *Stack) Push(value int) {
	for !s.q1.IsEmpty() {
		s.q2.Enqueue(s.q1.Dequeue())
	}
	s.q1.Enqueue(value)
	for !s.q2.IsEmpty() {
		s.q1.Enqueue(s.q2.Dequeue())
	}
}
func (s *Stack) IsEmpty() bool {
	return s.q1.front == nil
}
func (s *Stack) Pop() int {
	if s.q1.IsEmpty() {
		return -1
	}
	return s.q1.Dequeue()
}
func (s *Stack) Print() {
	for !s.q1.IsEmpty() {
		current := s.q1.Dequeue()
		fmt.Printf("%d ", current)
		s.q2.Enqueue(current)
	}
	for !s.q2.IsEmpty() {
		s.q1.Enqueue(s.q2.Dequeue())
	}
}
func main() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Print()
}
