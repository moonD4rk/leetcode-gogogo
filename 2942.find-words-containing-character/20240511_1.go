package main

func findWordsContaining(words []string, x byte) []int {
	var result []int
	for index, word := range words {
		for i := range word {
			// use this code not type transfer
			if word[i] == x {
				result = append(result, index)
				break
			}
		}
	}
	return result
}
