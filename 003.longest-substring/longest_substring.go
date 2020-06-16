package main

import "fmt"

func lengthOfLongestSubstring(s string) (result int) {
	var (
		start   int
		tempMap = map[rune]int{}
	)
	// pwwkew
	for i, v := range s {
		fmt.Println(i, string(v), string(tempMap[v]))
		if now, ok := tempMap[v]; ok && start <= now {
			start++
		} else {
			result = max(result, i-start+1)
		}
		tempMap[v] = i
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	strList := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	for _, v := range strList {
		l := lengthOfLongestSubstring(v)
		fmt.Println(l)
	}
}
