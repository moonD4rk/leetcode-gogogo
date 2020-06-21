package main

import "fmt"

var numbers = []int{
	1, 3, 5, 6,
}
var index = []int{
	5,
}

func searchInsert(nums []int, target int) int {
	var (
		result int
		mid    int
		length = len(nums) - 1
	)

	if target > nums[length] {
		return length + 1
	} else if target < nums[result] {
		return result
	} else {
		for result < length {
			// 找出中间值，二分法每次都找中间值
			mid = length + result / 2
			// 如果目标值等于中间值
			fmt.Println(target, result, length, mid)
			if target == nums[mid] {
				return mid
				// 如果目标小于中间值,则说明index在前半部分，将长度 -1 后继续二分
			} else if target > nums[mid] {
				result++
			} else {
				length--
			}
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
