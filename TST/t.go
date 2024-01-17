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
func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key}
	} else if root.key < key {
		root.right = insert(root.right, key)
	} else {
		root.left = insert(root.left, key)
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
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
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
		InorderTraversal(root.right)
		fmt.Printf(" %d ", root.key)
	}
}
func main() {
	t := Tree{}
	arr := []int{1, 22, 33, 44, 55, 66, -22, -44}
	for _, num := range arr {
		t.Insert(num)
	}
	InorderTraversal(t.root)
	fmt.Println()
	fmt.Println(t.Search(-44))
	t.Delete(-44)
	fmt.Println(t.Search(-44))
}
