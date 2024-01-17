package main

import "fmt"

func quickSort(arr []int, low, high int) {
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
	quickSort(arr, s, high)
	quickSort(arr, low, e)
}
func main() {
	arr := []int{1, 13, 30, -10, -18, 22, 21, 11}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
