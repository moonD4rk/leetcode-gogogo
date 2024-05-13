## Minimum Window Substring

Given two strings `s` and `t`, return *the minimum window in `s` which will contain all the characters in `t`*. If there is no such window in `s` that covers all characters in `t`, return *the empty string `""`*.

**Note** that If there is such a window, it is guaranteed that there will always be only one unique minimum window in `s`.

**Example 1:**

```
Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
```

**Example 2:**

```
Input: s = "a", t = "a"
Output: "a"
```

 

**Constraints:**

- `1 <= s.length, t.length <= 105`
- `s` and `t` consist of English letters.



## 解题思路

题目给定了两个字符串 S 和 T ，让你在第一个字符串中找出包含第二个串的最小子字符串。

### 思路一：暴力求解

1. 遍历字符串 S，从中找出T字符出现的位置，依次遍历直到最终答案。





