package main

import "fmt"

var numbers = []int{
	1, 3, 5, 6,
}
var index = []int{
	5, 0, 2, 7,
}

func searchInsert(nums []int, target int) int {
	var (
		result int
		length = len(nums) - 1
		mid    int
	)
	switch {
	case target > nums[length]:
		return length + 1
	case target < nums[result]:
		return result
	default:
		for result < length {
			mid = (result + length) / 2
			switch {
			case target < nums[mid]:
				length = mid
			case target == nums[mid]:
				return mid
			case target > nums[mid]:
				result = mid + 1
			}
		}
		return result
	}
}

func main() {
	for _, v := range index {
		result := searchInsert(numbers, v)
		fmt.Println(result)
	}
}
