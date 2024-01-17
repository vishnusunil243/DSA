package main

import "fmt"

type Heap struct {
	arr []int
}

func (h *Heap) insert(value int) {
	h.arr = append(h.arr, value)
	h.heapifyUp(len(h.arr) - 1)
}
func (h *Heap) heapifyUp(index int) {
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}
func (h *Heap) Delete() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.heapifyDown(0)
	return extracted
}
func HeapSort(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		swap(arr, i, 0)
		heapify(arr, i, 0)
	}

}
func heapify(arr []int, n, i int) {
	largest := i
	l, r := left(i), right(i)
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		fmt.Println(arr[largest], arr[i])
		swap(arr, largest, i)
		fmt.Println(arr[largest], arr[i])
		heapify(arr, n, largest)
	}
}
func (h *Heap) heapifyDown(index int) {
	lastIndex := len(h.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] > h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.arr[index] < h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			break
		}
	}
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
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
func main() {
	h := Heap{}
	arr := []int{12, 22, 202, 303, 55, 85, 505}
	for _, n := range arr {
		h.insert(n)
		fmt.Println(h)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(h.Delete())
		fmt.Println(h)
	}
	HeapSort(arr)
	fmt.Println(arr)
}
