## Move Zeroes

Given an array `nums`, write a function to move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Example:**

```
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
```

**Note**:

1. You must do this **in-place** without making a copy of the array.
2. Minimize the total number of operations.

### 解决方法

设定一个计数器 `count` ，遍历列表，当列表中的数不为 0 时，将当前值移动到 `nums[count]` ，`count`值加1，最终`count`为当前所有不为 0 的总数。

遍历第二个列表，从 `count` 位置开始，将 `nums` 后面的值设为 0. 

```go
package main

import "fmt"

var numbers = [][]int{
	{0, 1, 0, 3, 12},
	{0, 0, 0, 0, 12},
	{0, 0, 0, 0, 0},
}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	for _, v := range numbers {
		fmt.Println("before", v)
		moveZeroes(v)
		fmt.Println("after", v)
	}
}

```



