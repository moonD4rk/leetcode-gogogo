## Remove Duplicates from Sorted Array

Given a sorted array *nums*, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that each element appear only *once* and return the new length.

Do not allocate extra space for another array, you must do this by **modifying the input array [in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** with O(1) extra memory.

**Example 1:**

```
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
```

**Example 2:**

```
Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
```

**Clarification:**

Confused why the returned value is an integer but your answer is an array?

Note that the input array is passed in by **reference**, which means modification to the input array will be known to the caller as well.

Internally you can think of this:

```
// nums is passed in by reference. (i.e., without making a copy)
int len = removeDuplicates(nums);

// any modification to nums in your function would be known by the caller.
// using the length returned by your function, it prints the first len elements.
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
```

### 解决办法

题目要求空间复杂度为O(1)，所以不能使用其他数据结构

由于是有序列表，遍历列表，判断当前值和列表中前一个值是否相等便可。

```go
package main

import "fmt"

var numbers = [][]int{
	{1, 1, 2},
	{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	{1, 1},
	{1, 1, 1},
	{1, 2},
	{},
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var temp = 0
	for i, _ := range nums {
		if nums[i] != nums[temp] {
			temp++
			nums[temp] = nums[i]
		}
	}
	return temp + 1
}

func main() {
	for _, v := range numbers {
		lenNumbs := removeDuplicates(v)
		fmt.Println(lenNumbs)
	}
}
```

