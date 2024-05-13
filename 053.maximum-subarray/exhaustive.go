package main

import (
	"log"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			maxSum = max(maxSum, sum)
		}
	}
	return maxSum
}

func main() {
	s := []int{-2, -1}
	i := maxSubArray(s)
	log.Println(i)
}
