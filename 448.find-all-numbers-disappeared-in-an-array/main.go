package main

import (
	"fmt"
	"math"
)

var numbers = [][]int{
	{4, 3, 2, 7, 8, 2, 3, 1},
}

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0, len(nums))
	for _, num := range nums {
		i := int(math.Abs(float64(num))) - 1
		if nums[i] > 0 {
			nums[i] *= -1
		}
	}
	for i, num := range nums {
		if num > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	for _, v := range numbers {
		h := findDisappearedNumbers(v)
		fmt.Println(h)
	}
}
