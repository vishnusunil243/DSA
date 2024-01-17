package main

import "fmt"

const ArraySize = 7

type Hashtable struct {
	array [ArraySize]*Bucket
}
type Bucket struct {
	head *Node
}
type Node struct {
	key  string
	next *Node
}

func (h *Hashtable) insert(value string) {
	index := Hash(value)
	h.array[index].insert(value)
}
func (b *Bucket) insert(value string) {
	newNode := &Node{key: value}
	if b.head == nil {
		b.head = newNode
		return
	}
	current := b.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}
func (h *Hashtable) delete(value string) {
	index := Hash(value)
	fmt.Println(index)
	h.array[index].delete(value)
}
func (b *Bucket) delete(value string) {
	if b.head == nil {
		return
	}
	if b.head.key == value {
		b.head = b.head.next
		return
	}
	current := b.head
	for current.next != nil {
		if current.next.key == value {
			current.next = current.next.next
			fmt.Println("hi")
			return
		}
		current = current.next
	}

}
func (h *Hashtable) Search(value string) bool {
	index := Hash(value)
	return h.array[index].Search(value)
}
func (b *Bucket) Search(value string) bool {
	current := b.head
	for current != nil {
		if current.key == value {
			return true
		}
		current = current.next
	}
	return false
}
func Init() *Hashtable {
	H := Hashtable{}
	for i := range H.array {
		H.array[i] = &Bucket{}
	}
	return &H
}
func Hash(key string) int {
	res := 0
	for i := 0; i < len(key); i++ {
		res += int(key[i])
	}
	return res % ArraySize
}
func main() {
	h := Init()
	h.insert("VISHNU")
	h.insert("AKSHAY")
	h.delete("VISHNU")

	fmt.Println(h.Search("VISHNU"))
}
