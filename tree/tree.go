package main

import (
	"fmt"
	"math"
)

type Node struct {
	key   int
	left  *Node
	right *Node
}
type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value int) {
	t.Root = insert(t.Root, value)
}
func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{key: key}
	}
	if root.key < key {
		root.right = insert(root.right, key)
	} else if root.key > key {
		root.left = insert(root.left, key)
	}
	return root
}
func (t *Tree) Search(value int) bool {
	return search(t.Root, value)
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
func (t *Tree) Delete(value int) {
	t.Root = delete(t.Root, value)
}
func delete(root *Node, value int) *Node {
	if root.key > value {
		root.left = delete(root.left, value)
	} else if root.key < value {
		root.right = delete(root.right, value)
	} else {
		if root.right == nil {
			return root.left
		} else if root.left == nil {
			return root.right
		}
		root.key = inorderSuccessor(root.right)
		root.right = delete(root.right, value)
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
func InorderTraversal(root *Node) {
	if root != nil {
		InorderTraversal(root.left)
		fmt.Printf("%d => ", root.key)
		InorderTraversal(root.right)
	}
}
func PreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d => ", root.key)
		PreOrderTraversal(root.left)
		PreOrderTraversal(root.right)
	}
}
func PostOrderTraversal(root *Node) {
	if root != nil {
		PostOrderTraversal(root.left)
		PostOrderTraversal(root.right)
		fmt.Printf("%d => ", root.key)
	}
}
func closest(root *Node, target float64) int {
	closest := math.MaxInt64
	for root != nil {
		key := float64(root.key)
		if math.Abs(key-target) < math.Abs(float64(closest)-target) {
			closest = root.key
		}
		if key == target {
			break
		} else if key < target {
			root = root.right
		} else {
			root = root.left
		}
	}
	return closest
}
func IsValidBst(node *Node) bool {
	return isValidBst(node, math.MinInt64, math.MaxInt64)
}
func isValidBst(node *Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.key <= min || node.key >= max {
		return false
	}
	return isValidBst(node.left, min, node.key) && isValidBst(node.right, node.key, max)
}
func main() {
	t := Tree{}
	arr := []int{100, 18, 22, 23, 44, 16, 101}
	for _, num := range arr {
		t.Insert(num)
	}
	fmt.Println(closest(t.Root, 101))
	fmt.Println(IsValidBst(t.Root))

}
