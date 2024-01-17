package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*Bucket
}
type Bucket struct {
	head *Node
}
type Node struct {
	key  string
	next *Node
}

func (h *HashTable) insert(value string) {
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
func (h *HashTable) delete(value string) {
	index := Hash(value)
	h.array[index].delete(value)
}
func (b *Bucket) delete(value string) {
	if b.head.key == value {
		b.head = b.head.next
		return
	}
	current := b.head
	for current.next != nil {
		if current.next.key == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
func (h *HashTable) Search(value string) bool {
	index := Hash(value)
	return h.array[index].Search(value)
}
func (b *Bucket) Search(value string) bool {
	if b.head == nil {
		return false
	}
	current := b.head
	for current != nil {
		if current.key == value {
			return true
		}
		current = current.next
	}
	return false
}
func Hash(key string) int {
	res := 0
	for i := 0; i < len(key); i++ {
		res += int(key[i])
	}
	return res % ArraySize
}
func Init() *HashTable {
	h := HashTable{}
	for i := range h.array {
		h.array[i] = &Bucket{}
	}
	return &h
}
func main() {
	h := Init()
	h.insert("VISHNU")
	h.insert("RANDY")
	h.delete("VISHNU")
	fmt.Println(h.Search("VISHNU"))
}
