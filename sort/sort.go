package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}
}
func insertionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		last := len(arr) - i - 1
		max := getMax(arr, last)
		arr[max], arr[last] = arr[last], arr[max]
	}
}
func getMax(arr []int, last int) int {
	max := 0
	for i := 0; i <= last; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}
func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	s := low
	e := high
	m := s + (e-s)/2
	piv := arr[m]
	for s <= e {
		for arr[s] < piv {
			s++
		}
		for arr[e] > piv {
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
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)

}
func merge(left, right []int) []int {
	i := 0
	k := 0
	j := 0
	MergedArray := make([]int, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			MergedArray[k] = left[i]
			i++
		} else {
			MergedArray[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		MergedArray[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		MergedArray[k] = right[j]
		j++
		k++
	}
	return MergedArray
}
func main() {
	arr := []int{4, 5, -8, 3, -7}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
