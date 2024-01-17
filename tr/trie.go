package main

import "fmt"

type Trie struct {
	root *Node
}
type Node struct {
	children map[rune]*Node
	isEnd    bool
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
	deleteRecursive(s, t.root, 0)
}
func deleteRecursive(s string, root *Node, depth int) {
	if root == nil {
		return
	}
	if len(s) == depth {
		root.isEnd = false
	} else {
		char := rune(s[depth])
		nextNode, exists := root.children[char]
		if exists {
			deleteRecursive(s, nextNode, depth+1)
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
	fmt.Println(t.Search("VISHNUVY"))
}
