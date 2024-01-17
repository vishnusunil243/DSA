package main

import "fmt"

func binarySearch(arr []int, s, e, target int) int {
	if s > e {
		return -1
	}
	m := s + (e-s)/2
	if arr[m] == target {
		return m
	} else if arr[m] > target {
		return binarySearch(arr, 0, m-1, target)
	}
	return binarySearch(arr, m+1, e, target)
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 4))
}
