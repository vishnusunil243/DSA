package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) insert(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}
func (h *maxHeap) heapifyUp(index int) {
	for h.arr[parent(index)] > h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func (h *maxHeap) extract() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1
	if len(h.arr) == 0 {
		fmt.Println("cannot extract from an empty heap")
		return -1
	}
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.heapifyDown(0)
	return extracted
}
func (h *maxHeap) heapifyDown(index int) {
	lastIndex := len(h.arr) - 1
	l, r := leftChild(index), rightChild(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] < h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.arr[index] > h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftChild(index), rightChild(index)
		} else {
			break
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return (2 * i) + 1
}
func rightChild(i int) int {
	return (2 * i) + 2
}
func (h *maxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func main() {
	h := maxHeap{}
	arr := []int{1, 22, 43, 55}
	for _, i := range arr {
		h.insert(i)
		fmt.Println(h)
	}
	fmt.Println(h.extract())
	fmt.Println(h)
}
