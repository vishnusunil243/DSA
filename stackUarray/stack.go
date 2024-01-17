package main

import "fmt"

type Stack struct {
	items []int
}
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}
func main() {
	s := Stack{}
	s.Push(11)
	s.Push(1221)
	s.Pop()
	fmt.Println(s.items)
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(123)
	fmt.Println(q.items)
	q.Dequeue()
	fmt.Println(q.items)
}
