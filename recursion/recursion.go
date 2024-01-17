package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
func isSorted(arr []int, s int) bool {
	if s == len(arr)-1 {
		return true
	}
	if arr[s] <= arr[s+1] {
		return isSorted(arr, s+1)
	}
	return false
}
func linearSearch(arr []int, target, s int) int {
	if s == len(arr) {
		return -1
	}
	if arr[s] == target {
		return s
	}
	return linearSearch(arr, target, s+1)
}
func plusOne(digits []int) []int {
	var n string
	if digits[len(digits)-1] != 9 {
		digits[len(digits)-1]++
		return digits
	}
	for _, num := range digits {
		n += strconv.Itoa(num)
	}
	fmt.Println(n)
	s, _ := strconv.Atoi(n)
	fmt.Println(s)
	s++
	n = strconv.Itoa(s)
	var res []int
	for i := 0; i < len(n); i++ {
		r, _ := strconv.Atoi(string(n[i]))
		res = append(res, r)
	}
	return res

}
func main() {
	// fmt.Println(fibonacci(4))
	arr := []int{5, 2, 2, 6, 5, 7, 1, 9, 0, 3, 8, 6, 8, 6, 5, 2, 1, 8, 7, 9, 8, 3, 8, 4, 7, 2, 5, 8, 9}
	// fmt.Println(isSorted(arr, 0))
	// fmt.Println(linearSearch(arr, 99, 0))
	fmt.Println(plusOne(arr))

}
