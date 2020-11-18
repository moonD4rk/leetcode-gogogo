# longest_substring

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```


示例 2:

```
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```


示例 3:

```
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

### 解题

#### 方法一 暴力破解法
1. 首先两层 for 循环找出所有的子字符串
2. 再加一层 for 循环遍历子字符串，使用 map 判断子字符串中是否含有重复元素
```go
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
// Status: Time Limit Exceeded
```
结果是超出了规定的时间限制

#### 方法二：滑动窗口算法(sliding window)



所有子字符串的问题都应该用滑动窗口算法（sliding window）

打算所有子字符串题目放一起解决，此处先暂时放放。