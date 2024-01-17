package main

import "fmt"

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		swapped := false
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
func selection(arr []int) {
	for i := 0; i < len(arr); i++ {
		last := len(arr) - i - 1
		max := getMax(arr, last)
		arr[max], arr[last] = arr[last], arr[max]
	}
}
func getMax(arr []int, end int) int {
	max := 0
	for i := 0; i <= end; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}
func insertion(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}
func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	mergedArray := make([]int, len(left)+len(right))
	i := 0
	j := 0
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			mergedArray[k] = right[j]
			j++
		} else {
			mergedArray[k] = left[i]
			i++
		}
		k++
	}
	for i < len(left) {
		mergedArray[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		mergedArray[k] = right[j]
		j++
		k++
	}
	return mergedArray
}
func mergeSortInPlace(arr []int, s, e int) {
	if e-s == 1 {
		return
	}
	mid := (s + e) / 2
	mergeSortInPlace(arr, s, mid)
	mergeSortInPlace(arr, mid, e)
	mergeInPlace(arr, s, mid, e)
}

func mergeInPlace(arr []int, s, mid, e int) {
	mergedArray := make([]int, e-s)
	i := s
	j := mid
	k := 0
	for i < mid && j < e {
		if arr[i] > arr[j] {
			mergedArray[k] = arr[j]
			j++
		} else {
			mergedArray[k] = arr[i]
			i++
		}
		k++
	}
	for i < mid {
		mergedArray[k] = arr[i]
		i++
		k++
	}
	for j < e {
		mergedArray[k] = arr[j]
		j++
		k++
	}
	for l := 0; l < len(mergedArray); l++ {
		arr[s+l] = mergedArray[l]
	}
}
func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	s := low
	e := high
	m := s + (e-s)/2
	pivot := arr[m]
	for s <= e {
		for arr[s] < pivot {
			s++
		}
		for arr[e] > pivot {
			e--
		}
		if s <= e {
			arr[s], arr[e] = arr[e], arr[s]
			s++
			e--
		}
	}
	quickSort(arr, low, e)
	quickSort(arr, s, high)

}
func main() {
	arr := []int{1, 2, 3, 44, -88, -1, 22}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

}
