package main

import (
	"log"
)

var numbers = [][]int{
	{2, 2, 1},
	{4, 1, 2, 1, 2},
	{4, 1, 2, 1, 2, 4, 3},
}

func singleNumber(nums []int) int {
	var result int
	for _, v := range nums {
		result ^= v
	}
	return result
}

func main() {
	var (
		a = 10
		b = 9
	)
	c := a ^ b
	log.Println(c)

	for _, v := range numbers {
		r := singleNumber(v)
		log.Println(r)
	}
}
