package main

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	s := low
	e := high
	m := s + (e-s)/2
	p := arr[m]
	for s <= e {
		for arr[s] < p {
			s++
		}
		for arr[e] > p {
			e--
		}
		if s <= e {
			arr[s], arr[e] = arr[e], arr[s]
			e--
			s++
		}
	}
	QuickSort(arr, low, e)
	QuickSort(arr, s, high)
}
func main() {
	arr := []int{1, 2, -2, 3, -4}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
