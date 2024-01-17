package main

import "fmt"

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
		swap(arr, largest, i)
		heapify(arr, n, largest)
	}
}
func HeapSort(arr []int, n int) []int {
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		swap(arr, i, 0)
		heapify(arr, i, 0)
	}
	return arr
}
func left(i int) int {
	return (2 * i) + 1
}
func right(i int) int {
	return (2 * i) + 2
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	arr := []int{1, 2, 23, -99, 77, 54}
	fmt.Println(HeapSort(arr, len(arr)))
}
