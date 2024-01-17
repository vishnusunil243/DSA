package main

import "fmt"

type Node struct {
	children map[rune]*Node
	isEnd    bool
}
type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node)}}
}
func (t *Trie) Insert(s string) {
	node := t.root
	for _, char := range s {
		if node.children[char] == nil {
			node.children[char] = &Node{children: make(map[rune]*Node)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}
func (t *Trie) Delete(s string) {
	deleteRecursive(t.root, s, 0)
}
func deleteRecursive(node *Node, s string, depth int) {
	if node == nil {
		return
	}
	if len(s) == depth {
		node.isEnd = false
		return
	} else {
		char := rune(s[depth])
		nextNode, exists := node.children[char]
		if exists {
			deleteRecursive(nextNode, s, depth+1)
			if len(nextNode.children) == 0 && !nextNode.isEnd {
				delete(nextNode.children, char)
			}
		}
	}
}
func (t *Trie) Search(s string) bool {
	node := t.root
	for _, char := range s {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}
func main() {
	t := NewTrie()
	t.Insert("VISHNU")
	t.Insert("VISHNUVY")
	t.Delete("VISHNU")
	fmt.Println(t.Search("VISHNU"))
}
