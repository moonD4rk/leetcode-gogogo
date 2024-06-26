package main

import (
	"log"
)

func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, len(nums))
	for i, v := range nums {
		if _, ok := tempMap[v]; ok {
			return []int{tempMap[v], i}
		}
		tempMap[target-v] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	i := twoSum(nums, target)
	log.Println(i)
}
