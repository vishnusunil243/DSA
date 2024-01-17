package main

import "fmt"

func BinarySearch(arr []int, target, s, e int) int {
	if s > e {
		return -1
	}
	m := s + (e-s)/2
	if arr[m] == target {
		return m
	} else if arr[m] < target {
		return BinarySearch(arr, target, s, m-1)
	}

	return BinarySearch(arr, target, m+1, e)

}
func main() {
	arr := []int{1, 2, 3, 44, 55, 66}
	s := BinarySearch(arr, 3, 0, len(arr)-1)
	fmt.Println(s)
}
