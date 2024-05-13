package main

import (
	"log"
)

func main() {
	testCases := []struct {
		words  []string
		x      string
		wanted []int
	}{
		{
			words:  []string{"leet", "code"},
			x:      "e",
			wanted: []int{0, 1},
		},
		{
			words:  []string{"abc", "bcd", "aaaa", "cbc"},
			x:      "a",
			wanted: []int{0, 2},
		},
		{
			words:  []string{"abc", "bcd", "aaaa", "cbc"},
			x:      "z",
			wanted: []int{},
		},
	}

	for _, tc := range testCases {
		wanted := findWordsContaining(tc.words, tc.x[0])
		log.Println(wanted)
	}
}
