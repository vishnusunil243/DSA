package main

import "fmt"

type maxHeap struct {
	arr []int
}

func (h *maxHeap) insert(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}
func (h *maxHeap) heapifyUp(i int) {
	for h.arr[parent(i)] < h.arr[i] {
		h.swap(i, parent(i))
		i = parent(i)
	}
}
func (h *maxHeap) extract() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1
	if len(h.arr) == 0 {
		fmt.Println("cannot delete from an empty heap")
		return -1
	}
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.heapifyDown(0)
	return extracted
}
func (h *maxHeap) heapifyDown(i int) {
	lastIndex := len(h.arr) - 1
	l, r := left(i), right(i)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] > h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.arr[childToCompare] > h.arr[i] {
			h.swap(childToCompare, i)
			i = childToCompare
			l, r = left(i), right(r)
		} else {
			break
		}
	}
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}
func (h *maxHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func main() {
	h := maxHeap{}
	arr := []int{23, 33, 44, -88, 101, 202}
	for _, i := range arr {
		h.insert(i)
		fmt.Println(h)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(h.extract())
		fmt.Println(h)
	}
}
