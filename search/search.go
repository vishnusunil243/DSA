package main

import "fmt"

func binarySearch(arr []int, target, s, e int) int {
	if s > e {
		return -1
	}
	m := s + (e-s)/2
	if arr[m] == target {
		return m
	} else if arr[m] < target {
		return binarySearch(arr, target, m+1, e)
	}
	return binarySearch(arr, target, s, m-1)

}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := binarySearch(arr, 2, 0, len(arr)-1)
	fmt.Printf("%d ", res)
}
