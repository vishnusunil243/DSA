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
		return 0
	}
	toRemove := s.top.data
	s.top = s.top.next
	return toRemove
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
		current := q.Dequeue()
		fmt.Printf("%d ", current)
		temp.Enqueue(current)
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (s *Stack) StackToQueue() *Queue {
	temp := Stack{}
	res := Queue{}
	for !s.IsEmpty() {
		current := s.Pop()
		temp.Push(current)
	}
	for !temp.IsEmpty() {
		res.Enqueue(temp.Pop())
	}
	return &res
}
func (s *Stack) Print() {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		fmt.Printf("%d ", current)
		temp.Push(current)
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (q *Queue) insertAfter(value, target int) {
	if q.IsEmpty() {
		return
	}
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		temp.Enqueue(current)
		if current == target {
			temp.Enqueue(value)
		}
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (q *Queue) DeleteWithValue(value int) {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		if current != value {
			temp.Enqueue(current)
		}
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
func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(12)
	q.Enqueue(22)
	q.insertAfter(101, 12)
	q.DeleteWithValue(12)
	q.Print()
	fmt.Println()
	s := q.QueueToStack()
	s.Print()
}
