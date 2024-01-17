package main

import "fmt"

type Stack struct {
	arr []int
}

func (s *Stack) Push(value int) {
	s.arr = append(s.arr, value)
}
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	toRemove := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return toRemove
}
func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}
func (s *Stack) Print() {
	tmp := Stack{}
	for !s.IsEmpty() {
		current := s.Pop()
		fmt.Printf("%d ", current)
		tmp.Push(current)
	}
	for !tmp.IsEmpty() {
		s.Push(tmp.Pop())
	}
}
func main() {
	s := Stack{}
	s.Push(1)
	s.Push(12)
	s.Print()
}
