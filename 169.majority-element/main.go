package main

import "fmt"

var numbers = [][]int{
	{3, 2, 3},
	{2, 2, 1, 1, 1, 2, 2},
	{3, 3, 4},
	{6, 6, 6, 7, 7},
	{3, 1, 1, 1, 1, 4, 1, 2, 3, 4, 1, 1, 2},
}

func majorityElement(nums []int) int {
	var (
		count, temp int
	)
	for i := 0; i < len(nums); i++ {
		switch {
		case count == 0:
			temp = nums[i]
			count++
		case temp == nums[i]:
			count++
		default:
			count--
		}
	}
	return temp
}

func main() {
	for _, v := range numbers {
		s := majorityElement(v)
		fmt.Println(s)
	}
}
