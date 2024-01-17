package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (l *linkedlist) insertAtEnd(value int) {
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
func (l *linkedlist) reverse(node *node) {
	if node == nil {
		return
	}
	l.reverse(node.next)
	fmt.Printf("%d ", node.data)
}
func (l *linkedlist) insertAfter(value, target int) {
	newNode := &node{data: value}
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
	newNode := &node{data: value}
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
func (l linkedlist) print() {
	if l.head == nil {
		return
	}
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
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
func (l *linkedlist) Check() bool {
	if l.head == nil {
		return false
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}
func binarySearch(arr []int, value, s, e int) int {
	if s > e {
		return -1
	}
	m := s + (e-s)/2
	if arr[m] == value {
		return m
	} else if arr[m] < value {
		return binarySearch(arr, value, m+1, e)
	}
	return binarySearch(arr, value, s, m-1)

}
func (l *linkedlist) FindMiddle() int {
	if l.head == nil {
		return -1
	}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.data

}
func ArrayToLinkedList(arr []int) *linkedlist {
	l := linkedlist{}
	for _, num := range arr {
		l.insertAtEnd(num)
	}
	return &l
}
func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[indices[i]]
	}
	var r string
	for _, str := range res {
		r += string(str)
	}
	return r
}
func main() {
	// list := linkedlist{}
	// list.insertAtEnd(1)
	// list.insertAtEnd(2)
	// list.insertAtEnd(3)
	// list.print()
	// fmt.Println()
	// list.reverse(list.head)
	// list.insertAfter(22, 2)
	// list.insertAfter(11, 3)
	// list.insertBefore(12, 11)
	// list.DeleteWithValue(0)
	// fmt.Println()
	// list.print()
	// fmt.Println()
	// arr := []int{1, 2, 3, 4, 5, 7}
	// // fmt.Println(binarySearch(arr, 5, 0, len(arr)-1))
	// l := ArrayToLinkedList(arr)
	// l.print()
	// fmt.Printf("the middle element is %d", l.FindMiddle())
	// l.head.next.next = l.head.next
	// fmt.Println(l.Check())
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))

	// [4,5,6,7,0,2,1,3]
}
