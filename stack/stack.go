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
	newNode := &Node{data: value, next: s.top}
	s.top = newNode
}
func (s *Stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	topData := s.top.data
	s.top = s.top.next
	return topData
}
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.top.data
}
func (s *Stack) Size() int {
	count := 0
	current := s.top
	for current != nil {
		count++
		current = current.next
	}
	return count
}
func (s *Stack) insertAfter(target, value int) {
	temp := Stack{}
	for !s.IsEmpty() {
		if s.Peek() == target {
			temp.Push(value)
			break
		}
		current := s.Pop()
		temp.Push(current)

	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func Print(node *Node) {
	if node == nil {
		return
	}
	Print(node.next)
	fmt.Printf("%d ", node.data)
}
func (s *Stack) insertBefore(value, target int) {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		temp.Push(current)
		if current == target {
			temp.Push(value)
		}
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (s *Stack) DeleteWithValue(value int) {
	temp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		if current != value {
			temp.Push(current)
		}
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (s *Stack) Reverse() {
	if s.top == nil {
		return
	}
	current := s.top
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(33)

	stack.insertAfter(2, 13)
	stack.insertBefore(22, 2)
	stack.DeleteWithValue(1)
	stack.Reverse()

}
