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
		current := q.Dequeue()
		fmt.Printf("%d ", current)
		temp.Enqueue(current)
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
func (q *Queue) QueueToStack() *Stack {
	temp := &Stack{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		temp.Push(current)
	}
	return temp
}

// func StringReverse(s string) string {
// 	temp := Stack{}
// 	var res []byte
// 	for i := 0; i < len(s); i++ {
// 		temp.Push(s[i])
// 	}
// 	for !temp.IsEmpty() {
// 		res = append(res, temp.Pop())
// 	}
// 	return string(res)
// }
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
func (s *Stack) StackToQueue() *Queue {
	temp := Stack{}
	q := Queue{}
	for !s.IsEmpty() {
		current := s.Pop()
		temp.Push(current)
	}
	for !temp.IsEmpty() {
		current := temp.Pop()
		q.Enqueue(current)
		// s.Push(current)
	}
	return &q
}
func main() {
	q := Queue{}
	q.Enqueue(12)
	q.Enqueue(22)
	q.Enqueue(24)
	q.insertAfter(25, 12)
	q.Print()
}
