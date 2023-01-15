package main

import (
	"fmt"
)

var (
	revSymbols = map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}
	s1    = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	s10   = []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	s100  = []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	s1000 = []string{"M", "MM", "MMM"}
)

func intToRoman(i int) string {
	var (
		k1000, k100, k10, k int
	)
	k1000 = i / 1000
	result := ""
	if k1000 > 0 {
		result += s1000[k1000-1]
		i = i - k1000*1000
	}
	k100 = i / 100
	if k100 > 0 {
		result += s100[k100-1]
		i = i - k100*100
	}
	k10 = i / 10
	if k10 > 0 {
		result += s10[k10-1]
		i = i - k10*10
	}
	k = i
	if k > 0 {
		result += s1[k-1]
	}
	return result
}

func main() {
	n := intToRoman(18)
	fmt.Println(n)
}
