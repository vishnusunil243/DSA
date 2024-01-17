package main

import "fmt"

type node struct {
	data int
	next *node
	prev *node
}
type dll struct {
	head *node
	tail *node
}

func (dl *dll) insert(value int) {
	newNode := &node{data: value}
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
		return
	}
	current := dl.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	dl.tail = newNode
	newNode.prev = current
}
func (dl *dll) print() {
	if dl.head == nil {
		return
	}
	current := dl.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func (dl *dll) Reverse() {
	if dl.tail == nil {
		return
	}
	current := dl.tail
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.prev
	}
}
func (dl *dll) insertAfter(value, target int) {
	newNode := &node{data: value}
	if dl.head == nil {
		return
	}

	if dl.tail.data == target {
		dl.tail.next = newNode
		newNode.prev = dl.tail
		dl.tail = newNode
		return
	}
	current := dl.head
	for current != nil {
		if current.data == target {
			newNode.next = current.next
			current.next = newNode
			newNode.prev = current
			newNode.next.prev = newNode
			return
		}
		current = current.next
	}
}
func (dl *dll) insertBefore(value, target int) {
	newNode := &node{data: value}
	if dl.head.data == target {
		newNode.next = dl.head
		dl.head.prev = newNode
		dl.head = newNode
		return
	}
	current := dl.head
	for current.next != nil {
		if current.next.data == target {
			newNode.next = current.next
			newNode.prev = current
			newNode.next.prev = newNode
			current.next = newNode
			return
		}
		current = current.next
	}
}
func (dl *dll) DeleteWithValue(value int) {
	if dl.head == nil {
		return
	}
	if dl.head.data == value {
		dl.head = dl.head.next
		dl.head.prev = nil
		return
	}
	if dl.tail.data == value {
		dl.tail = dl.tail.prev
		dl.tail.next = nil
		return
	}
	current := dl.head
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
	dll := dll{}
	dll.insert(1)
	dll.insert(2)
	dll.insert(44)
	// dll.print()
	// dll.Reverse()
	dll.insertAfter(11, 1)
	dll.insertBefore(22, 2)
	dll.DeleteWithValue(2)
	dll.print()
	fmt.Println()
	dll.Reverse()

}
