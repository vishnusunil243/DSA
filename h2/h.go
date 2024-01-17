package main

import "fmt"

type Heap struct {
	array []int
}

func (h *Heap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}
func (h *Heap) heapifyUp(index int) {
	for h.array[Parent(index)] < h.array[index] {
		h.Swap(Parent(index), index)
		index = Parent(index)
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
	l, r := Left(index), Right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.Swap(index, childToCompare)
			index = childToCompare
			l, r = Left(index), Right(index)
		} else {
			break

		}
	}
}
func HeapSort(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		swap(arr, i, 0)
		heapify(arr, i, 0)
	}
}
func heapify(arr []int, n, i int) {
	largest := i
	l, r := Left(i), Right(i)
	if l < n && arr[largest] < arr[l] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}
	if largest != i {
		swap(arr, largest, i)
		heapify(arr, n, largest)
	}
}
func Parent(i int) int {
	return (i - 1) / 2
}
func Left(i int) int {
	return (2 * i) + 1
}
func Right(i int) int {
	return (2 * i) + 2
}
func (h *Heap) Swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	arr := []int{12, 22, 44, 55, 909, -111}
	HeapSort(arr)
	fmt.Println(arr)
}
