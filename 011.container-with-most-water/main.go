package main

import "log"

var numbers = [][]int{
	{1, 8, 6, 2, 5, 4, 8, 3, 7},
	{1, 2, 3, 4, 5},
	{0, 0},
	{1, 3},
	{1, 1},
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	if len(height) == 2 {
		return short(height[0], height[1])
	}
	var area int
	for i := 0; i < len(height); i++ {
		var temp int
		for j := i; j < len(height); j++ {
			high := short(height[i], height[j])
			length := j - i
			temp = high * length
			if temp > area {
				area = temp
			}
		}
	}
	return area
}

func short(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxArea2(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var area int
	left := 0
	right := len(height) - 1
	for left <= right {
		tempArea := (right - left) * short(height[left], height[right])
		if tempArea > area {
			area = tempArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}

func main() {
	for _, v := range numbers {
		s := maxArea2(v)
		log.Println(s)
	}
}
