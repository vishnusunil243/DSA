package main

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	var res int64
	for len(nums) != 1 {
		first := strconv.Itoa(nums[0])
		second := strconv.Itoa(nums[len(nums)-1])
		total := first + second
		totalVal, _ := strconv.Atoi(total)
		fmt.Println(totalVal)
		res += int64(totalVal)
		fmt.Println(res)
		nums = nums[1 : len(nums)-1]
	}
	if len(nums) == 1 {
		res += int64(nums[0])
	}
	return res
}
func main() {
	arr := []int{5, 14, 13, 8, 12}
	fmt.Println(findTheArrayConcVal(arr))
}
