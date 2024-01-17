package main

import (
	"fmt"
	"math"
)

type Node struct {
	data  int
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
		return &Node{data: value}
	} else if root.data < value {
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
	} else if root.data < value {
		return search(root.right, value)
	} else if root.data > value {
		return search(root.left, value)
	}
	return true
}
func (t *Tree) Delete(value int) {
	t.root = delete(t.root, value)
}
func delete(root *Node, value int) *Node {
	if root == nil {
		return nil
	} else if root.data < value {
		root.right = delete(root.right, value)
	} else if root.data > value {
		root.left = delete(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {
			root.data = inorderSuccessor(root.right)
			root.right = delete(root.right, root.data)
		}
	}
	return root
}
func inorderSuccessor(root *Node) int {
	current := root
	for current.left != nil {
		current = current.left
	}
	return current.data
}
func closest(root *Node, value float64) int {
	closest := math.MaxInt64
	for root != nil {
		key := float64(root.data)
		if math.Abs(key-value) < math.Abs(float64(closest)-value) {
			closest = root.data
		}
		if key == value {
			break
		} else if root.data < int(value) {
			root = root.right
		} else {
			root = root.left
		}
	}
	return closest
}
func (t *Tree) IsValidBst() bool {
	return isValidBst(t.root, math.MaxInt64, math.MinInt64)
}
func isValidBst(node *Node, max, min int) bool {
	if node == nil {
		return true
	}
	if node.data <= min || node.data >= max {
		return false
	}
	return isValidBst(node.left, node.data, min) && isValidBst(node.right, max, node.data)
}
func main() {
	t := Tree{}
	t.Insert(12)
	t.Insert(14)
	t.Insert(13)
	t.Insert(22)
	t.Insert(10)
	fmt.Println(t.Search(12))
	t.Delete(12)
	fmt.Println(closest(t.root, 20))
	fmt.Println(t.IsValidBst())
}
