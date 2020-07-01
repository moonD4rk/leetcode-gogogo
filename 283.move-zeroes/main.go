package main

import "fmt"

var numbers = [][]int{
	{0, 1, 0, 3, 12},
	{0, 0, 0, 0, 12},
	{0, 0, 0, 0, 0},
}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	for _, v := range numbers {
		fmt.Println("before", v)
		moveZeroes(v)
		fmt.Println("after", v)
	}
}
