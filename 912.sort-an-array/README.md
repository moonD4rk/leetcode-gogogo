## sort-an-array

Given an array of integers `nums`, sort the array in ascending order.

**Example 1:**

```
Input: nums = [5,2,3,1]
Output: [1,2,3,5]
```

**Example 2:**

```
Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
```

**Constraints:**

- `1 <= nums.length <= 50000`
- `-50000 <= nums[i] <= 50000`

### 解题思路

#### 冒泡法

一次比较两个元素，如果他们的顺序错误就把他们交换过来，时间复杂度O(n^2)

```go
func sortBubble(nums []int) []int {
   for i := 0; i < len(nums); i++ {
      for j := i; j < len(nums); j++ {
         if nums[i] >= nums[j] {
            nums[i], nums[j] = nums[j], nums[i]
            continue
         }
      }
   }
   return nums
}
// Runtime: 3004 ms, faster than 5.90% of Go online submissions for Sort an Array.
// Memory Usage: 6.5 MB, less than 59.04% of Go online submissions for Sort an Array.
```

#### 快速排序

时间复杂度 O(nlog(n))，最差为 O(n2)

快速排序使用[分治法](https://zh.wikipedia.org/wiki/分治法)（Divide and conquer）策略来把一个[序列](https://zh.wikipedia.org/wiki/序列)（list）分为较小和较大的2个子序列，然后递归地排序两个子序列。

1. 挑选基准值：从数列中挑出一个元素，称为“基准”（pivot）
2. 分割：重新排序数列，所有比基准值小的元素摆放在基准前面，所有比基准值大的元素摆在基准后面（与基准值相等的数可以到任何一边）。在这个分割结束之后，对基准值的排序就已经完成
3. 递归排序子序列：[递归](https://zh.wikipedia.org/wiki/递归)地将小于基准值元素的子序列和大于基准值元素的子序列排序。

```go
package main

import (
	"fmt"
)

func quickSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	// 选择一个基准值
	pivot := len(nums) / 2
	start := 0
	end := len(nums) - 1
	// 将基准值放到最后
	fmt.Println(nums[pivot])
	nums[pivot], nums[end] = nums[end], nums[pivot]
	fmt.Println(nums)
	for i, _ := range nums {
		if nums[i] < nums[end] {
			fmt.Printf("%d < %d, 将当前值 %d 和标记值 %d 交换位置\n", nums[i], nums[end], nums[i], nums[start])
			nums[i], nums[start] = nums[start], nums[i]
			fmt.Println(nums)
			start++
		} else {
			fmt.Printf("%d > %d, 队列不变\n", nums[i], nums[end])
			fmt.Println(nums)
		}
	}
	fmt.Println(nums)
	// 遍历完成后，再将基准值换回到标记位置
	// 此标记位置为队列中最后一个小于基准值的位置
	fmt.Println(start, nums[start])
	nums[start], nums[end] = nums[end], nums[end]
	// 此时，已根据基准值将数组分成了两组
	// 小于基准值的在队列左边，大于基准值的在队列右边
	// 将两个队列不停的分为大于基准值和小于基准值的两个子队列
	quickSort(nums[start+1:])
	quickSort(nums[:start])
	return nums
}
```

#### 逻辑

```
随机选择一个基准值 4
[6 2 1 5 9 7 3 8 10 1 2 3 4]
6 > 4，不变
[6 2 1 5 9 7 3 8 10 1 2 3 4]
2 < 4，
[2 6 1 5 9 7 3 8 10 1 2 3 4]
1 < 4
[2 1 6 5 9 7 3 8 10 1 2 3 4]
5 > 4
[2 1 6 5 9 7 3 8 10 1 2 3 4]
9 > 4
[2 1 6 5 9 7 3 8 10 1 2 3 4]
7 > 4
[2 1 6 5 9 7 3 8 10 1 2 3 4]
3 < 4
[2 1 3 5 9 7 6 8 10 1 2 3 4]
8 > 4
[2 1 3 5 9 7 6 8 10 1 2 3 4]
10 > 4
[2 1 3 5 9 7 6 8 10 1 2 3 4]
1 < 4
[2 1 3 1 9 7 6 8 10 5 2 3 4]
2 < 4
[2 1 3 1 2 7 6 8 10 5 9 3 4]
3 < 4
[2 1 3 1 2 3 6 8 10 5 9 7 4]
4 > 4
[2 1 3 1 2 3 6 8 10 5 9 7 4]
将基准值交换回标记位置，4和6交换
[2 1 3 1 2 3 4 8 10 5 9 7 6]
此时，再将队列分为两个部分，分别选择各自的基准值进行排序
[2 1 3 1 2 3 4] 
[8 10 5 9 7 6]

```



