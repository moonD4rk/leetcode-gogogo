package main

import "fmt"

var numbers = []int{
	1, 3, 5, 6,
}
var index = []int{
	5, 2, 7, 0,
}

func searchInsert(nums []int, target int) int {
	var result int
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i, v := range nums {
		if target <= v {
			result = i
			break
		}
	}
	return result
}

func main() {
	for _, v := range index {
		result := searchInsert(numbers, v)
		fmt.Println(result)
	}
}
