package main

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}
func (h *Heap) heapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
func (h *Heap) Delete() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.heapifyDown(0)
	return extracted
}
func (h *Heap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), righ(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[childToCompare] > h.array[index] {
			h.swap(childToCompare, index)
			index = childToCompare
			l, r = left(index), righ(index)
		} else {
			break
		}
	}
}
func (h *Heap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return (2 * i) + 1
}
func righ(i int) int {
	return (2 * i) + 2
}
func main() {
	h := Heap{}
	arr := []int{12, 2, 3, 4, 22}
	for _, i := range arr {
		h.insert(i)
		fmt.Println(h)
	}
	for i := 0; i < 3; i++ {
		h.Delete()
		fmt.Println(h)
	}
}
