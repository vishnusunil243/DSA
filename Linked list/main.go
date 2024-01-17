package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l linkedList) Print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
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
	l.length++

}
func (l *linkedList) insertAfterTarget(targetValue, value int) {
	newNode := &node{data: value}
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		if current.data == targetValue {
			newNode.next = current.next
			current.next = newNode
			l.length++
		}
		current = current.next
	}
}
func (l *linkedList) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	previousToDelete := l.head
	if value == l.head.data {
		l.head = l.head.next
		l.length--
		return
	}
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
func (l *linkedList) insertBeforeTarget(targetValue, value int) {
	newNode := &node{data: value}
	if l.head == nil {
		l.prepend(newNode)
		return
	}
	if targetValue == l.head.data {
		newNode.next = l.head
		l.head = newNode
		l.length++
		return
	}

	current := l.head
	for current.next != nil {

		if current.next.data == targetValue {
			newNode.next = current.next
			current.next = newNode
			l.length++
			return
		}
		current = current.next
	}
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 77}
	node2 := &node{data: 74}
	node3 := &node{data: 99}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.DeleteWithValue(74)
	myList.insertAtEnd(74)
	myList.insertAfterTarget(77, 61)
	myList.insertBeforeTarget(77, 16)
	myList.Print()

}
