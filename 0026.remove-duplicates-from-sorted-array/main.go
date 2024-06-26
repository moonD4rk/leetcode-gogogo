package main

import (
	"log"
)

var numbers = [][]int{
	{1, 1, 2},
	{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	{1, 1},
	{1, 1, 1},
	{1, 2},
	{},
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	temp := 0
	for i := range nums {
		if nums[i] != nums[temp] {
			temp++
			nums[temp] = nums[i]
		}
	}
	return temp + 1
}

func main() {
	for _, v := range numbers {
		lenNumbs := removeDuplicates(v)
		log.Println(lenNumbs)
	}
}
