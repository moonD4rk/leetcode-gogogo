package main

import (
	"fmt"
)

var allNums = [][]int{
	{6, 2, 1, 5, 9, 7, 4, 8, 10, 1, 2, 3, 3},
	// {5, 2, 3, 1},
	// {5, 1, 1, 2, 0, 0},
	// {5, 4, 3, 2, 1, 0},
	// {5, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0},
	// {5, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0, 4, 3, 2, 1, 0},
}

func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	// 选择一个基准值
	pivot := len(nums) / 2
	start := 0
	end := len(nums) - 1
	// 将基准值放到最后
	fmt.Println(nums[pivot])
	nums[pivot], nums[end] = nums[end], nums[pivot]
	fmt.Println(nums)
	for i, _ := range nums {
		if nums[i] < nums[end] {
			fmt.Printf("%d < %d, 将当前值 %d 和标记值 %d 交换位置\n", nums[i], nums[end], nums[i], nums[start])
			nums[i], nums[start] = nums[start], nums[i]
			fmt.Println(nums)
			start++
		} else {
			fmt.Printf("%d > %d, 队列不变\n", nums[i], nums[end])
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)
	// 遍历完成后，再将基准值换回到标记位置
	// 此标记位置为队列中最后一个小于基准值的位置
	fmt.Println(start, nums[start])
	nums[start], nums[end] = nums[end], nums[end]
	// 此时，已根据基准值将数组分成了两组
	// 小于基准值的在队列左边，大于基准值的在队列右边
	// 将两个队列不停的分为大于基准值和小于基准值的两个子队列
	quickSort(nums[start+1:])
	quickSort(nums[:start])
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
		s := quickSort(v)
		fmt.Println(s)
	}
}
