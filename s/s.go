package main

import "fmt"

type Node struct {
	data byte
	next *Node
}
type Stack struct {
	top *Node
}

func (s *Stack) Push(value byte) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return '0'
	}
	toRemove := s.top.data
	s.top = s.top.next
	return toRemove
}

//	func Paranthesis(s string) bool {
//		stack := Stack{}
//		for i := 0; i < len(s); i++ {
//			if s[i] == '(' {
//				stack.Push(s[i])
//			} else if s[i] == ')' && stack.Pop() != '(' {
//				return false
//			}
//		}
//		if !stack.IsEmpty() {
//			return false
//		}
//		return true
//	}
// type Queue struct {
// 	s1, s2 *Stack
// }

//	func NewQueue() *Queue {
//		return &Queue{
//			s1: &Stack{},
//			s2: &Stack{},
//		}
//	}
//
//	func (q *Queue) Enqueue(value int) {
//		for !q.s1.IsEmpty() {
//			q.s2.Push(q.s1.Pop())
//		}
//		q.s1.Push(value)
//		for !q.s2.IsEmpty() {
//			q.s1.Push(q.s2.Pop())
//		}
//	}
//
//	func (q *Queue) Dequeue() int {
//		return q.s1.Pop()
//	}
//
//	func (q *Queue) IsEmpty() bool {
//		return q.s1.IsEmpty()
//	}
func StringReverse(s string) string {
	tmp := Stack{}
	var res string
	for i := 0; i < len(s); i++ {
		tmp.Push(s[i])
	}
	for !tmp.IsEmpty() {
		res += string(tmp.Pop())
	}
	return res
}
func main() {
	fmt.Println(StringReverse("vishnu"))
}
