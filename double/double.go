package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}
type linkedList struct {
	head *Node
}

func (l *linkedList) insertAtEnd(value int) {
	newNode := &Node{data: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current

}
func (l *linkedList) Print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (l *linkedList) Reverse() {
	node := l.head
	for node.next != nil {
		node = node.next
	}

	for node != nil {
		fmt.Printf("%d ", node.data)
		node = node.prev
	}
}
func (l *linkedList) insertAfter(value, target int) {
	newNode := &Node{data: value}
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		if current.data == target {
			newNode.next = current.next
			newNode.prev = current
			current.next = newNode
			return
		}
		current = current.next
	}
}
func (l *linkedList) insertBefore(value, target int) {
	newNode := &Node{data: value}
	if l.head == nil {
		return
	}
	if l.head.data == target {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == target {
			newNode.next = current.next
			newNode.prev = current
			current.next = newNode
			newNode.next.prev = newNode
			return
		}
		current = current.next
	}
}
func (l *linkedList) DeleteWithValue(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.head.prev = nil
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			current.next.prev = current
			return
		}
		current = current.next
	}
}
func main() {
	list := linkedList{}
	list.insertAtEnd(1)
	list.insertAtEnd(3)
	list.Print()
	fmt.Println()
	list.Reverse()
	fmt.Println()
	list.insertAfter(33, 3)
	list.Reverse()
	fmt.Println()
	list.insertBefore(11, 3)
	list.Print()
	fmt.Println()
	list.Reverse()
	fmt.Println()
	list.DeleteWithValue(11)
	list.Reverse()
}
