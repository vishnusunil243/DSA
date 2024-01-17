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
func (t *Trie) insert(s string) {
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
	t.deleteRecursive(t.root, s, 0)
}
func (t *Trie) deleteRecursive(node *Node, key string, depth int) {
	if node == nil {
		return
	}
	if len(key) == depth {
		node.isEnd = false
	} else {
		char := rune(key[depth])
		nextNode, exists := node.children[char]
		if exists {
			t.deleteRecursive(nextNode, key, depth+1)
			if len(nextNode.children) == 0 && !nextNode.isEnd {
				delete(node.children, char)
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
func (t *Trie) SearchPrefix(s string) bool {
	node := t.root
	for _, char := range s {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return true
}
func (t *Trie) PrintElementsWithPrefix(s string) []string {
	result := make([]string, 0)
	node := t.findNodeByPrefix(s)
	t.collectElementsWithPrefix(node, s, &result)
	return result
}
func (t *Trie) collectElementsWithPrefix(node *Node, prefix string, result *[]string) {
	if node == nil {
		return
	}
	if node.isEnd {
		*result = append(*result, prefix)
	}
	for char, child := range node.children {
		t.collectElementsWithPrefix(child, prefix+string(char), *&result)
	}
}
func (t *Trie) findNodeByPrefix(prefix string) *Node {
	node := t.root
	for _, char := range prefix {
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}
func main() {
	t := NewTrie()
	t.insert("VISHNU")
	t.insert("VISHHY")
	t.insert("VISN")
	result := t.PrintElementsWithPrefix("VI")
	fmt.Println(result)

}
