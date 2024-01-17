package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}
func Merge(left, right []int) []int {
	i := 0
	j := 0
	k := 0
	res := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		j++
		k++
	}
	return res
}
func main() {
	arr := []int{1, -11, 12, 22, 14}
	arr = MergeSort(arr)
	fmt.Println(arr)
}
