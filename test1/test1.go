package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type linkedlist struct {
	head *Node
}

func (l *linkedlist) insertAtEnd(value int) {
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
}
func (l linkedlist) Print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (l *linkedlist) insertAfter(value, target int) {
	newNode := &Node{data: value}
	if l.head == nil {
		return
	}
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
func (l *linkedlist) insertBefore(value, target int) {
	newNode := &Node{data: value}
	if l.head == nil {
		return
	}
	if l.head.data == target {
		newNode.next = l.head
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == target {
			newNode.next = current.next
			current.next = newNode
			return
		}
		current = current.next
	}
}
func (l *linkedlist) DeleteWithValue(value int) {
	if l.head == nil {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		return
	}
	current := l.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
func Reverse(node *Node) {
	if node == nil {
		return
	}
	Reverse(node.next)
	fmt.Printf("%d ", node.data)

}
func (l *linkedlist) Check() bool {
	if l.head == nil {
		return false
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			return true
		}
	}
	return false

}
func (l *linkedlist) Search(value int) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	count := 1
	for current != nil {
		if current.data == value {
			return count
		}
		count++
		current = current.next

	}
	return -1
}
func main() {
	list := linkedlist{}
	list.insertAtEnd(1)
	list.insertAtEnd(2)
	list.insertAtEnd(5)
	list.Print()
	fmt.Println()
	Reverse(list.head)
	fmt.Println()
	list.insertAfter(44, 5)
	list.Print()
	fmt.Println()
	list.insertBefore(11, 44)
	list.Print()
	fmt.Println()
	list.DeleteWithValue(1)
	list.Print()
	fmt.Println()
	fmt.Printf("%d ", list.Search(44))
	fmt.Println()
	list.head.next.next = list.head.next
	fmt.Println(list.Check())
}
