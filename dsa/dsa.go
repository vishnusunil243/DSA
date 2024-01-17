package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
}

func (l *linkedList) insertAtEnd(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (l *linkedList) prepend(value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode

}
func (l *linkedList) insertAfter(insertValue, target int) {
	if l.head == nil {
		return
	}
	newNode := &node{data: insertValue}
	current := l.head
	for current != nil {
		if current.data == target {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}
func (l *linkedList) DeleteWithValue(value int) {
	if l.head == nil {
		return
	}
	current := l.head
	if value == l.head.data {
		l.head = current.next
		return
	}
	for current != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
func (l *linkedList) insertBefore(value, target int) {
	if l.head == nil {
		return
	}
	newNode := &node{data: value}
	current := l.head
	if target == current.data {
		newNode.next = l.head
		l.head = newNode
		return
	}
	for current.next != nil {
		if target == current.next.data {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}
func (l *linkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (l *linkedList) middle() int {
	if l.head == nil {
		return -1
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	return slow.data
}
func (l *linkedList) Reverse(node *node) {
	if node == nil {
		return
	}
	l.Reverse(node.next)
	fmt.Printf("%d ", node.data)
}
func main() {
	list := &linkedList{}
	list.prepend(2)
	list.prepend(3)
	list.prepend(9)
	list.insertAfter(16, 2)
	list.insertBefore(18, 3)
	// list.DeleteWithValue(9)

	fmt.Println(list.middle())
	list.print()
	// list.Reverse(list.head)
}
