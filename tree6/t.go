package main

import "fmt"

type Node struct {
	key   int
	left  *Node
	right *Node
}
type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	t.root = insert(t.root, value)
}
func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{key: value}
	}
	if root.key < value {
		root.right = insert(root.right, value)
	} else {
		root.left = insert(root.left, value)
	}
	return root
}
func (t *Tree) Delete(value int) {
	t.root = delete(t.root, value)
}
func delete(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.key < value {
		root.right = delete(root.right, value)
	} else if root.key > value {
		root.left = delete(root.left, value)
	} else {
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		root.key = inorderSuccessor(root.right)
		root.right = delete(root.right, root.key)
	}
	return root
}
func inorderSuccessor(root *Node) int {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current.key
}
func (t *Tree) Search(value int) bool {
	return search(t.root, value)
}
func search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.key < value {
		return search(root.right, value)
	} else if root.key > value {
		return search(root.left, value)
	}
	return true
}
func InorderTraversal(root *Node) {
	if root != nil {
		InorderTraversal(root.left)
		fmt.Printf(" %d ", root.key)
		InorderTraversal(root.right)
	}
}
func main() {
	arr := []int{12, 22, 44, 55, 77, 10}
	t := Tree{}
	for _, num := range arr {
		t.Insert(num)
	}
	t.Delete(44)
	InorderTraversal(t.root)
}
