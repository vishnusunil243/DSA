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
func Closest(root *Node, value float64) int {
	closest := math.MaxInt64
	for root != nil {
		key := float64(root.key)
		if math.Abs(key-value) < math.Abs(float64(closest)-value) {
			closest = root.key
		}
		if key == value {
			break
		} else if key < value {
			root = root.right
		} else {
			root = root.left
		}
	}
	return closest
}
func (t *Tree) IsValidBst() bool {
	return isValidBst(t.root, math.MinInt64, math.MaxInt64)
}
func isValidBst(root *Node, min, max int) bool {
	if root == nil {
		return true
	}
	if root.key <= min || root.key >= max {
		return false
	}
	return isValidBst(root.left, min, root.key) && isValidBst(root.right, root.key, max)
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
	arr := []int{12, 22, 33, 44, -88, -90, 10}
	for _, n := range arr {
		t.Insert(n)
	}
	fmt.Println(t.Search(22))
	t.Delete(22)
	fmt.Println(t.Search(22))
	InorderTraversal(t.root)
	fmt.Println()
	fmt.Println(t.IsValidBst())
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
	fmt.Println(s.IsValidBst())
	fmt.Println(Closest(t.root, -80))
}
