package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(value int) {
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
	return q.front == nil
}
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	toRemove := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return toRemove
}
func (q *Queue) Print() {
	tmp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		fmt.Printf("%d ", current)
		tmp.Enqueue(current)
	}
	for !tmp.IsEmpty() {
		q.Enqueue(tmp.Dequeue())
	}
}
func main() {
	q := Queue{}
	q.Enqueue(12)
	q.Enqueue(22)
	q.Print()
}
