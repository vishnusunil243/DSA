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
	newNode := &Node{data: value, next: nil}
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
	frontData := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	return frontData
}
func (q *Queue) Size() int {
	count := 0
	current := q.front
	for current != nil {
		count++
		current = current.next
	}
	return count
}
func (q *Queue) insertAfter(value, target int) {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		temp.Enqueue(current)
		if current == target {
			temp.Enqueue(value)
		}
	}

	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (q *Queue) insertBefore(value, target int) {
	temp := Queue{}

	for !q.IsEmpty() {
		current := q.Dequeue()
		if current == target {
			temp.Enqueue(value)
		}
		temp.Enqueue(current)
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (q *Queue) DeleteWithValue(value int) {
	temp := Queue{}
	for !q.IsEmpty() {
		current := q.Dequeue()
		if current != value {
			temp.Enqueue(current)
		}
	}
	for !temp.IsEmpty() {
		q.Enqueue(temp.Dequeue())
	}
}
func (q *Queue) Print() {
	current := q.front
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
}
func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(22)
	q.Enqueue(23)
	q.insertAfter(11, 22)
	q.insertBefore(20, 22)
	q.DeleteWithValue(1)
	q.Print()

}
