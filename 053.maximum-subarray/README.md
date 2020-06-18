# 53. Maximum Subarray

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
```
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:
```


If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
## 解决方法

### 枚举法
枚举所有可能性，找出其中最大的值，比如说列表长度为 N，则最终的子列表数目为 A(N,N-1)
时间复杂度为O(N^2)
```go
func maxSubArray(nums []int) int {
	maxSum := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			maxSum = max(maxSum, sum)
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
```

### 拆分法

将列表从中拆成两个，则最大子列表的值可能在左边、右边、和中间三种可能性。

分别计算左边和右边，如果不在左边或右边，出从中间序列的值中，分别向左右两边依次扩展。