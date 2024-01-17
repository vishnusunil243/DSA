package main

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
}

func (l *linkedlist) append(value int) {
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
func (l *linkedlist) prepend(value int) {
	newNode := &node{data: value}
	newNode.next = l.head
	l.head = newNode
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
func Reverse(s string) {
	for i := 0; i < len(s); i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
func binarySearch(arr []int, s, e, target int) int {
	if s > e {
		return -1
	}
	m := s + (e-s)/2
	if arr[m] == target {
		return m
	} else if arr[m] > target {
		return binarySearch(arr, s, m-1, target)
	}
	return binarySearch(arr, m+1, e, target)

}
