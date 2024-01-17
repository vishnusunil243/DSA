package main

import "fmt"

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

const ArraySize = 7

func (h *HashTable) insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
func (h *HashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].Delete(key)
}

func (b *Bucket) insert(k string) {
	newBucketNode := &Node{key: k}
	newBucketNode.next = b.head
	b.head = newBucketNode
}
func (b *Bucket) search(k string) bool {
	current := b.head
	for current != nil {
		if current.key == k {
			return true
		}
		current = current.next
	}
	return false
}
func (b *Bucket) Delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	current := b.head
	for current.next != nil {
		if current.next.key == k {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
func hash(key string) int {
	var res int
	for i := 0; i < len(key); i++ {
		res += int(key[i])
	}
	return res % ArraySize
}
func Init() *HashTable {
	newHashTable := HashTable{}
	for i := range newHashTable.array {
		newHashTable.array[i] = &Bucket{}
	}
	return &newHashTable
}
func main() {
	hashTaable := Init()
	hashTaable.insert("RANDY")
	hashTaable.insert("KANE")
	hashTaable.Delete("RANDY")
	fmt.Println(hashTaable.search("RANDY"))
}
