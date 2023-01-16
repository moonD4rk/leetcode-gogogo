package main

var stringList = [][]string{
	{"ADOBECODEBANC", "ABC"}, // BANC
	{"a", "a"},
}

func minWindow(s string, t string) string {
	var (
		tmap       = make(map[byte]int)
		tempString string
		count      = 0
	)
	for _, v := range []byte(t) {
		if tmap[v] == 0 {
			count++
		}
		tmap[v]++
	}

	return tempString
}

func main() {
	for _, v := range stringList {
		minWindow(v[0], v[1])
	}
}
