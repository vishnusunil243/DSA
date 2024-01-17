package main

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	l, r := left(i), right(i) // Fix: Use left(i) and right(i) instead of l and r
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
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func heapSort(arr []int, n int) []int {
	for i := n / 2; i >= 0; i-- { // Fix: Start from n/2 instead of (n/2)-1
		heapify(arr, n, i)
	}
	for i := n - 1; i > 0; i-- { // Fix: Loop should start from n-1
		swap(arr, 0, i)
		heapify(arr, i, 0) // Fix: Use i instead of n
	}
	return arr
}

func main() {
	arr := []int{12, -22, -44, 22, 24}
	fmt.Println(heapSort(arr, len(arr)))
}
