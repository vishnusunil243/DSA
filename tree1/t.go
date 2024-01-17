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
	}
	if root.data < value {
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
	if root.data < value {
		root.right = delete(root.right, value)
	} else if root.data > value {
		root.left = delete(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		root.data = inorderSuccessor(root.right)
		root.right = delete(root.right, root.data)
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
func (t *Tree) Search(value int) bool {
	return search(t.root, value)
}
func search(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if root.data < value {
		return search(root.right, value)
	} else if root.data > value {
		return search(root.left, value)
	}
	return true

}
func Closest(root *Node, value float64) int {
	closest := math.MaxInt64
	for root != nil {
		key := float64(root.data)
		if math.Abs(key-value) < math.Abs(float64(closest)-value) {
			closest = root.data
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
	return isValidBst(t.root, math.MaxInt64, math.MinInt64)
}
func isValidBst(root *Node, max, min int) bool {
	if root == nil {
		return true
	}
	if root.data <= min || root.data >= max {
		return false
	}
	return isValidBst(root.left, root.data, min) && isValidBst(root.right, max, root.data)
}

func main() {
	t := Tree{}
	arr := []int{12, 22, 44, 55, 10, 2, 3}
	for _, num := range arr {
		t.Insert(num)
	}
	fmt.Println(t.Search(22))
	t.Delete(22)
	fmt.Println(t.Search(22))
	fmt.Println(Closest(t.root, 50))
	fmt.Println(t.IsValidBst())
	s := Tree{}
	s.root = &Node{
		data: 10,
		left: &Node{
			data:  15, // Invalid: left child greater than root
			left:  &Node{data: 5},
			right: &Node{data: 12},
		},
		right: &Node{
			data:  20,
			left:  &Node{data: 18},
			right: &Node{data: 22},
		},
	}
	fmt.Println(s.IsValidBst())
}
