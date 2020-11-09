package main

import (
	"fmt"
	"strconv"
	"strings"
)

var versions = [][]string{
	{"0.1", "1.1"}, // 1
	{"1.0.1", "1"}, // -1
	{"7.5.2.4", "7.5.3"}, // 1
	{"7.5.2.4", "7.5.2.4.1"}, // 1
	{"1.01", "1.001"},  // 0
	{"1.0.0.0", "1.0"}, // 0
}

func main() {
	for _, v := range versions {
		i := compareVersion(v[0], v[1])
		fmt.Println(i)
	}
}

func compareVersion(version1 string, version2 string) int {
	v1Slice := strings.Split(version1, ".")
	v2Slice := strings.Split(version2, ".")
	for len(v1Slice) < len(v2Slice) {
		v1Slice = append(v1Slice, "0")
	}
	for len(v2Slice) < len(v1Slice) {
		v2Slice = append(v2Slice, "0")
	}
	mLength := minLength(v1Slice, v2Slice)
	for i := 0; i < mLength; i++ {
		v1Int, _ := strconv.Atoi(v1Slice[i])
		v2Int, _ := strconv.Atoi(v2Slice[i])
		if v1Int < v2Int {
			return -1
		} else if v1Int == v2Int {
			continue
		} else {
			return 1
		}
	}
	return 0
}

// minLength return min length between two slice
func minLength(a, b []string) int {
	if len(a) <= len(b) {
		return len(a)
	} else {
		return len(b)
	}
}

