package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type ListNode struct {
	head *Node
}

func (l *ListNode) Check() bool {
	if l.head == nil {
		return false
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false

}
func (l *ListNode) insertAtEnd(value int) {
	newNode := &Node{Value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (l *ListNode) print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
}
func (l *ListNode) DuplicateDelete() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil && current.Next != nil {
		if current.Value == current.Next.Value {
			current.Next = current.Next.Next
		}
		current = current.Next

	}
}

func main() {
	list := ListNode{}
	list.insertAtEnd(1)
	list.insertAtEnd(2)
	list.insertAtEnd(2)
	list.insertAtEnd(10)
	list.insertAtEnd(10)
	fmt.Println(list.Check())
	list.print()
	fmt.Println()
	list.DuplicateDelete()
	list.print()
	// list.head.Next.Next = list.head.Next
	// fmt.Println(list.Check())

}
