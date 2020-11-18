package main

import "fmt"

var subStrings = []string{
	"abcabcbb", // 3
	"bbbbb",    // 1
	"pwwkew",   // 3
	"",         // 0
}

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

// abcabcbb
func bruteSolution(s string) (result int) {
	var max int
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subString := s[i:j]
			if isUnique(subString) {
				if len(subString) > max {
					max = len(subString)
				}
			}
		}
	}
	return max
}

func isUnique(sub string) bool {
	var tempMap = make(map[int32]bool)
	for _, v := range sub {
		if _, ok := tempMap[v]; ok {
			return false
		}
		tempMap[v] = true
	}
	return true
}

func main() {
	for _, v := range subStrings {
		l := bruteSolution(v)
		fmt.Println(l)
	}
}
