## 35. Search Insert Position

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

**Example 1:**

```
Input: [1,3,5,6], 5
Output: 2
```

**Example 2:**

```
Input: [1,3,5,6], 2
Output: 1
```

**Example 3:**

```
Input: [1,3,5,6], 7
Output: 4
```

**Example 4:**

```
Input: [1,3,5,6], 0
Output: 0
```

### 解决方法

1. 遍历一遍数组，一次判断当前数和 target 的大小。

```go
func searchInsert(nums []int, target int) int {
	var result int
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i, v := range nums {
		if target <= v {
			result = i
			break
		}
	}
	return result
}
// Runtime: 4 ms, faster than 89.08% of Go online submissions for Search Insert Position.
// Memory Usage: 3.1 MB, less than 17.44% of Go online submissions for Search Insert Position.

```

