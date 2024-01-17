package main

import "fmt"

func Heapsort(arr []int) {
	n := len(arr)
	for i := n / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		swap(arr, 0, i)
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
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
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
	arr := []int{12, 22, 202, 502, -11}
	Heapsort(arr)
	fmt.Println(arr)
}
