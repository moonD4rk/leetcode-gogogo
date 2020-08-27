### Container With Most Water

[题目链接](https://leetcode.com/problems/trapping-rain-water/)

Given *n* non-negative integers *a1*, *a2*, ..., *an* , where each represents a point at coordinate (*i*, *ai*). *n* vertical lines are drawn such that the two endpoints of line *i* is at (*i*, *ai*) and (*i*, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

**Note:** You may not slant the container and *n* is at least 2.

![](https://darkmoon1973-cdn.oss-cn-beijing.aliyuncs.com/img/leetcode-11.png)

**Example:**

```
Input: [1,8,6,2,5,4,8,3,7]
Output: 49
```

### 解决办法

题目的意思是给你一个列表长度，让你求出列表长度和高度之间能成盛水的最大值。

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

##### 思路二：使用双指针

两个指针分别指向两端，计算两个指针之间的面积。如果要移动指针，一边指针不动，向中间移动指针指向值较小指针。最后等到两个指针相遇则说明计算完成，只用遍历一次数组，时间复杂度O(n)。

```go
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var area int
	left := 0
	right := len(height) - 1
	for left <= right {
		tempArea := (right - left) * short(height[left], height[right])
		if tempArea > area {
			area = tempArea
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
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
// Runtime: 12 ms, faster than 92.17% of Go online submissions for Container With Most Water.
// Memory Usage: 5.8 MB, less than 77.55% of Go online submissions for Container With Most Water.
```

这里通过移动指针来控制指针的边界，指针被移动则说明放弃此边界，最后完成求值。