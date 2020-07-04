## Find All Numbers Disappeared in an Array

Given an array of integers where 1 ≤ a[i] ≤ *n* (*n* = size of array), some elements appear twice and others appear once.

Find all the elements of [1, *n*] inclusive that do not appear in this array.

Could you do it without extra space and in O(*n*) runtime? You may assume the returned list does not count as extra space.

**Example:**

```
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
```



### 解决方法

最开始使用普通算法，时间复杂度超过了 `5.4%` 的人，查阅资料发现了一个巧妙的解法。遍历数组，通过标记正负数来传递额外信息。

遍历数组，将每一个元素对映的索引位置的值变为负数，遍历完成后，**空缺数对应的索引值则为正**，从而得出结果。

```
package main

import (
	"fmt"
	"math"
)

var numbers = [][]int{
	{4, 3, 2, 7, 8, 2, 3, 1},
}

func findDisappearedNumbers(nums []int) []int {
	ret := make([]int, 0, len(nums))
	for _, num := range nums {
		i := int(math.Abs(float64(num))) - 1
		if nums[i] > 0 {
			nums[i] *= -1
		}
	}
	for i, num := range nums {
		if num > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	for _, v := range numbers {
		h := findDisappearedNumbers(v)
		fmt.Println(h)
	}
}

```

