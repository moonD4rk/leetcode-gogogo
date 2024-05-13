package main

import (
	"log"
)

var stock = [][]int{
	{7, 1, 2, 0, 6, 4, 7}, // 6-1=5
	{7, 6, 4, 3, 1},       // 0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var maxInt int
	temp := prices[0]
	for _, v := range prices {
		if (v - temp) < 0 {
			temp = v
		} else {
			curr := v - temp
			maxInt = max(curr, maxInt)
		}
	}
	return maxInt
}

func main() {
	for _, v := range stock {
		log.Println(maxProfit(v))
	}
}
