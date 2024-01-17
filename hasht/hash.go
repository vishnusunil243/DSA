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
	data string
	next *Node
}

func (h *HashTable) insert(value string) {
	index := Hash(value)
	h.array[index].insert(value)
}
func (b *Bucket) insert(value string) {
	newNode := &Node{data: value}
	newNode.next = b.head
	b.head = newNode
}
func (h *HashTable) delete(value string) {
	index := Hash(value)
	h.array[index].delete(value)
}
func (h *HashTable) search(value string) bool {
	index := Hash(value)
	return h.array[index].search(value)
}
func (b *Bucket) search(value string) bool {
	current := b.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}
func (b *Bucket) delete(value string) {
	if b.head.data == value {
		b.head = b.head.next
		return
	}
	current := b.head
	for current.next != nil {
		if current.next.data == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}

}
func Hash(key string) int {
	res := 0
	for i := 0; i < len(key); i++ {
		res += int(key[i])
	}

	return res % ArraySize
}
func Init() *HashTable {
	H := HashTable{}
	for i := range H.array {
		H.array[i] = &Bucket{}
	}
	return &H

}
func main() {
	h := Init()
	h.insert("VISHNU")
	h.insert("VISHNU")
	h.insert("RAKESH")
	h.delete("RAKESH")
	fmt.Println(h.search("RAKESH"))
}
