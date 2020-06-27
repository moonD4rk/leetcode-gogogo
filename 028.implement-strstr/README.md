## Implement strStr().

Implement [strStr()](http://www.cplusplus.com/reference/cstring/strstr/).

Return the index of the first occurrence of needle in haystack, or **-1** if needle is not part of haystack.

**Example 1:**

```
Input: haystack = "hello", needle = "ll"
Output: 2
```

**Example 2:**

```
Input: haystack = "aaaaa", needle = "bba"
Output: -1
```


### 解决方法
1. 遍历主字符串，如果当前主字符(i+j)和子字符(j)不等，则 break 子字符串循环，当判断完结果相等后，满足给定条件 `len(needle) == j` 时，返回当前i值。
```go
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

```
