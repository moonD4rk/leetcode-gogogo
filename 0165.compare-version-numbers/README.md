## Compare Version Numbers

给定两个 version 版本，让你比较两个版本的大小

```
输入: version1 = "0.1", version2 = "1.1"
输出: -1

输入: version1 = "1.0.1", version2 = "1"
输出: 1


输入: version1 = "7.5.2.4", version2 = "7.5.3"
输出: -1

输入：version1 = "1.01", version2 = "1.001"
输出：0
解释：忽略前导零，“01” 和 “001” 表示相同的数字 “1”。
```

### 解题

首先将版本字符串切割为 Slice，然后比较 `Slice` 里面值的大小。

切记要对 `Slice` 补 0，将两个 `Slice` 的长度做成一致，否则会跑不过 `"1.0.0", "1.0"` 这条测试样例


``` go
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

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Compare Version Numbers.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Compare Version Numbers.
```


