package main

import "fmt"

var haystacks = [][]string{
	{"hello", "ll"},
	{"aaaaa", "aab"},
	{"fndainf1z", "nf1"},
	{"mississippi", "issipi"},
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}

func main() {
	for _, v := range haystacks {
		r := strStr(v[0], v[1])
		fmt.Println(r)
	}
}
