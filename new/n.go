package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Stack struct {
	top *Node
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

type Queue struct {
	s1, s2 *Stack
}

func NewQueue() *Queue {
	return &Queue{
		s1: &Stack{},
		s2: &Stack{},
	}
}
func (q *Queue) Enqueue(value int) {

	for !q.s2.IsEmpty() {
		current := q.s2.Pop()
		q.s1.Push(current)
	}
	q.s1.Push(value)
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
}
func (q *Queue) Dequeue() int {
	return q.s2.Pop()
}
func (q *Queue) Print() {
	for !q.s2.IsEmpty() {
		current := q.s2.Pop()
		fmt.Printf("%d ", current)
		q.s1.Push(current)
	}
	for !q.s1.IsEmpty() {
		q.s2.Push(q.s1.Pop())
	}
}
func main() {
	q := NewQueue()
	q.Enqueue(12)
	q.Enqueue(22)
	q.Print()
}
