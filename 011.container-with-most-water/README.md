### Container With Most Water

Given *n* non-negative integers *a1*, *a2*, ..., *an* , where each represents a point at coordinate (*i*, *ai*). *n* vertical lines are drawn such that the two endpoints of line *i* is at (*i*, *ai*) and (*i*, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

**Note:** You may not slant the container and *n* is at least 2.



**Example:**

```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

### 解决办法

题目的意思是给你一个列表长度，让你求出长度高度之间的最大值。

##### 思路一：全局遍历

遍历数组，求出当前面积的最大值，最后返回最大值

```go
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	if len(height) == 2 {
		return short(height[0], height[1])
	}
	var area int
	for i := 0; i < len(height); i++ {
		var temp int
		for j := i; j < len(height); j++ {
			high := short(height[i], height[j])
			length := j - i
			temp = high * length
			if temp > area {
				area = temp
			}
		}
	}
	return area
}

func short(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// Runtime: 316 ms, faster than 18.17% of Go online submissions for Container With Most Water.
// Memory Usage: 5.8 MB, less than 88.77% of Go online submissions for Container With Most Water.

```

可见时间复杂度是比较高的




