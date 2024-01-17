package main

import (
	"fmt"
)

type Node struct {
	data any
	next *Node
}
type Stack struct {
	top *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (s *Stack) Push(value any) {
	newNode := &Node{data: value}
	newNode.next = s.top
	s.top = newNode
}
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return -1
	}
	topData := s.top.data
	s.top = s.top.next
	return topData
}
func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return -1
	}
	return s.top.data
}
func (s *Stack) Print() {
	current := s.top
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (q *Queue) Enqueue(value any) {
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
	return q.rear == nil
}
func (q *Queue) Dequeue() any {
	if q.IsEmpty() {
		return -1
	}
	toRemove := q.front
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return toRemove.data
}
func (q *Queue) Print() {
	current := q.front
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (q *Queue) QueueToStack() *Stack {

	stack := Stack{}
	for !q.IsEmpty() {
		stack.Push(q.Dequeue())
	}
	return &stack
}
func (s *Stack) StackToQueue() *Queue {
	queue := Queue{}
	temp := Stack{}
	for !s.IsEmpty() {
		temp.Push(s.Pop())
	}
	for !temp.IsEmpty() {
		queue.Enqueue(temp.Pop())
	}
	return &queue
}
func (s *Stack) Size() int {
	if s.top == nil {
		return -1
	}
	count := 0
	temp := Stack{}
	for !s.IsEmpty() {
		temp.Push(s.Pop())
		count++
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
	return count
}
func (s *Stack) DeleteMiddle() any {
	if s.top == nil {
		return -1
	}
	temp := Stack{}
	mid := s.Size() / 2
	count := 0
	var toPop any
	for !s.IsEmpty() {
		if count == mid {
			toPop = s.Pop()
		} else {
			temp.Push(s.Pop())
		}
		count++
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
	return toPop
}
func (q *Queue) Size() int {
	count := 0
	temp := Queue{}
	for !q.IsEmpty() {
		temp.Enqueue(q.Dequeue())
		count++
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
	return count
}
func (s *Stack) insertMiddle(value any) {
	if s.IsEmpty() {
		s.Push(value)
		return
	}
	mid := s.Size() / 2
	count := 0
	temp := Stack{}
	for !s.IsEmpty() {
		if count == mid {
			temp.Push(value)
		}
		temp.Push(s.Pop())
		count++
	}
	for !temp.IsEmpty() {
		s.Push(temp.Pop())
	}
}
func (q *Queue) DeleteMiddle() any {
	if q.IsEmpty() {
		return -1
	}
	mid := q.Size() / 2
	count := 0
	temp := Queue{}
	var res any
	for !q.IsEmpty() {
		current := q.Dequeue()
		if count != mid {
			temp.Enqueue(current)
		} else {
			res = current
		}
		count++
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
	return res
}
func (q *Queue) insertMiddle(value any) {
	if q.IsEmpty() {
		q.Enqueue(value)
		return
	}
	temp := Queue{}
	mid := q.Size() / 2
	count := 0
	for !q.IsEmpty() {
		current := q.Dequeue()
		if count == mid {
			temp.Enqueue(value)
		}
		temp.Enqueue(current)
		count++
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func StringReverse(s string) string {
	temp := Stack{}
	res := ""
	for i := 0; i < len(s); i++ {
		temp.Push(s[i])
	}
	for !temp.IsEmpty() {
		res += string(temp.Pop().(byte))
	}
	return res
}
func isBalanced(s string) bool {
	temp := Stack{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			temp.Push(s[i])
		} else if s[i] == ')' {
			if temp.IsEmpty() || temp.Pop().(byte) != '(' {
				return false
			}
		}
	}
	if !temp.IsEmpty() {
		return false
	}
	return true
}
func main() {
	// s := Stack{}
	// s.Push(12)
	// s.Push(11)
	// s.Push(22)
	// s.insertMiddle(211)
	// // // fmt.Println(s.Pop())
	// // // fmt.Println(s.Pop())
	// // // fmt.Println(s.Peek())
	// // // fmt.Println(s.Pop())
	// // // fmt.Println(s.Pop())
	// fmt.Println(s.Size())
	// s.Print()
	// fmt.Println()
	// q := s.StackToQueue()
	// q.Print()
	// q := Queue{}
	// q.Enqueue(11)
	// q.Enqueue(12)
	// q.Enqueue(15)
	// // q.Enqueue(22)
	// q.insertMiddle(23)
	// q.Print()
	fmt.Println(isBalanced("(suku))"))
	// fmt.Println()
	// s := q.QueueToStack()
	// s.Print()
	// s := StringReverse("vishnu")
	// fmt.Printf("%s ", s)

}
