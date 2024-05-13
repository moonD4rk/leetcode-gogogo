package solution

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	digitLetters := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var result []string
	for _, v := range digits {
		letters := digitLetters[string(v)]
		if len(letters) == 0 {
			return []string{}
		} else {
			var next []string
			if len(result) == 0 {
				result = letters
			} else {
				for _, item2 := range result {
					for _, item1 := range letters {
						next = append(next, item2+item1)
					}
				}

				result = next
			}
		}
	}
	return result
}

// combine
func combine(raw [][]string) []string {
	if len(raw) == 0 {
		return []string{}
	}

	result := raw[0]

	for i := 1; i < len(raw); i++ {
		var next []string
		for _, item1 := range result {
			for _, item2 := range raw[i] {
				next = append(next, item1+item2)
			}
		}
		result = next
	}

	return result
}

func combine2(raw [][]string) []string {
	if len(raw) == 0 {
		return []string{}
	}

	result := raw[0]
	for _, s := range raw[1:] {
		var next []string
		for _, item1 := range result {
			for _, h2 := range s {
				next = append(next, item1+h2)
			}
		}
		result = next
	}
	return result
}

func combineThird(raw [][][]string) []string {
	if len(raw) == 0 {
		return []string{}
	}
	result2 := make([][]string, 0, len(raw))
	for _, sRaw := range raw {
		s := combine(sRaw)
		result2 = append(result2, s)
	}

	return combine(result2)
}

func combine3(raw [][]string) []string {
	if len(raw) == 0 {
		return []string{}
	}
	result := raw[0]
	for _, v := range raw[1:] {
		result = combineHandle(result, v)
	}
	return result
}

func combineHandle(raw1, raw2 []string) []string {
	result := make([]string, 0, len(raw1)*len(raw2))
	for _, r1 := range raw1 {
		for _, r2 := range raw2 {
			result = append(result, r1+r2)
		}
	}
	return result
}
