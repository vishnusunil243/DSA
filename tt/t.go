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
	root *Node
}

func (t *Tree) Insert(value int) {
	t.root = insert(t.root, value)
}
func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{key: value}
	} else if root.key < value {
		root.right = insert(root.right, value)
	} else {
		root.left = insert(root.left, value)
	}
	return root
}
func (t *Tree) Search(value int) bool {
	return search(t.root, value)
}
func search(root *Node, value int) bool {
	if root == nil {
		return false
	} else if root.key < value {
		return search(root.right, value)
	} else if root.key > value {
		return search(root.left, value)
	}
	return true
}
func (t *Tree) Delete(value int) {
	t.root = delete(t.root, value)
}
func delete(root *Node, value int) *Node {
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
func (t *Tree) closest(value float64) int {
	node := t.root
	closest := math.MaxInt64
	for node != nil {
		key := float64(node.key)
		if math.Abs(key-value) < math.Abs(float64(closest)-value) {
			closest = int(key)
		}
		if key == value {
			break
		} else if key < value {
			node = node.right
		} else {
			node = node.left
		}
	}
	return closest
}
func IsValidBst(node *Node) bool {
	return isValidBst(node, math.MaxInt64, math.MinInt64)
}
func isValidBst(node *Node, max, min int) bool {
	if node == nil {
		return true
	}
	if node.key <= min || node.key >= max {
		return false
	}
	return isValidBst(node.left, node.key, min) && isValidBst(node.right, max, node.key)
}
func inorderSuccessor(node *Node) int {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.key
}
func main() {
	t := Tree{}
	arr := []int{12, -12, 22, 24, 20}
	for _, num := range arr {
		t.Insert(num)
	}
	fmt.Println(IsValidBst(t.root))
	s := Tree{}
	s.root = &Node{
		key: 10,
		left: &Node{
			key:   15, // Invalid: left child greater than root
			left:  &Node{key: 5},
			right: &Node{key: 12},
		},
		right: &Node{
			key:   20,
			left:  &Node{key: 18},
			right: &Node{key: 22},
		},
	}
	fmt.Println(IsValidBst(s.root))

}
