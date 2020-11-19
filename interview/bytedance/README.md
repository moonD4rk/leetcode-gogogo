# 字符串移除

ATC?GA??，将 ? 替换为 ATCG 中的任何一个字符，
使得目标串ATCG的个数相同，输出目标串，如无解则输出===

遍历字符串，判断出每个正常字符有几个
如果总数除以几个正常字符的值不等于0，则返回===
如果没个数出现的个数大于需要出现的数，则返回===
比如 A 出现了3 次，需要出现的数 2

遍历字符串，遇到非正常字符替换为正常字符

## 解题
```go
func replaceString(s string, r string) string {
	uniqStr := make(map[string]int)
	for _, v := range s {
		if v != []rune(r)[0] {
			if _, ok := uniqStr[string(v)]; ok {
				uniqStr[string(v)]++
			} else {
				uniqStr[string(v)] = 1
			}
		}
	}
	need := len(s) % len(uniqStr)
	// h 为需要补齐的个数
	h := len(s) / len(uniqStr)
	for _, v := range uniqStr {
		if v > h {
			return "==="
		}
	}
	var result = ""
	if need == 0 {
		for _, v := range s {
			// 如果 v == ?, 则从 uniqStr 拿出小于 m 的字符串
			if v == []rune(r)[0] {
				for key, value := range uniqStr {
					if value < h {
						uniqStr[key]++
						result += key
						break
					}
				}
				continue
			}
			result += string(v)
		}
	} else {
		return "==="
	}
	return result
}

func main() {
	s := "ATC?"
	replace := "?"
	r := replaceString(s, replace)
	fmt.Println(r)
}
```