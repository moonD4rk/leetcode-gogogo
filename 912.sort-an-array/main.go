package main

import "fmt"

var allNums = [][]int{
	{5, 2, 3, 1},
	{5, 1, 1, 2, 0, 0},
	{5, 4, 3, 2, 1, 0},
	{5, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0},
	{5, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0},
}

func sortQuick(nums []int) []int {
	return nums
}

func sortBubble(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] >= nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	for _, v := range allNums {
		s := sortBubble(v)
		fmt.Println(s)
	}
}
