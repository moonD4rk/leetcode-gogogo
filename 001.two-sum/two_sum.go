package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var tempMap = make(map[int]int, len(nums))
	for i, v := range nums {
		tempMap[target-v] = i
	}
	for i, v := range nums {
		if _, ok := tempMap[v]; ok && tempMap[v] != i {
			return []int{i, tempMap[v]}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4, 4}
	target := 8
	i := twoSum(nums, target)
	fmt.Println(i)
}
