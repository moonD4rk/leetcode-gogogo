package main

import "fmt"

var stock = [][]int{
	{7, 1, 2, 0, 6, 4, 7}, // 6-1=5
	{7, 6, 4, 3, 1},       // 0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var max int
	temp := prices[0]
	for _, v := range prices {
		if (v - temp) < 0 {
			temp = v
		} else {
			curr := v - temp
			max = maxInt(curr, max)
		}
	}
	return max
}
func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	for _, v := range stock {
		max := maxProfit(v)
		fmt.Println(max)
	}

}
