## Majority Element

Given an array of size *n*, find the majority element. The majority element is the element that appears **more than** `⌊ n/2 ⌋` times.

You may assume that the array is non-empty and the majority element always exist in the array.

**Example 1:**

```
Input: [3,2,3]
Output: 3
```

**Example 2:**

```
Input: [2,2,1,1,1,2,2]
Output: 2
```

### 解题思路

要求找到列表里超过一半的数，列表不为空并且列表肯定存在一个超过一个超过一半的数。

1.遍历数组，将出现次数放到一个 map 中，如果值超过 len(n)/2 则返回。

```go
func majorityElement(nums []int) int {
    boudary := len(nums) / 2
    count := make(map[int]int)
    for _, v := range nums {
        count[v]++
        if count[v] > boudary { return v }
    }
    return -1
}
// Runtime: 20 ms, faster than 55.08% of Go online submissions for Majority Element.
//Memory Usage: 6 MB, less than 32.03% of Go online submissions for Majority Element.

```

emm…时间空间都不够理想

2.已知列表中肯定有数大于 n/2，那么此数出现的次数肯定要比其他数多。我们可以设定一个计数值 `count`，如果当前值和上个值相等时，`count` +1 ，如果不相等时 `count` 减1，**当 `count` 为 0 时，说明列表之前部分的计数中，前面值出现的次数为一样，前面所有计数都不算数，可以重新计数**

   ```go
package main

import "fmt"

var numbers = [][]int{
	{3, 2, 3},
	{2, 2, 1, 1, 1, 2, 2},
	{3, 3, 4},
	{6, 6, 6, 7, 7},
	{3, 1, 1, 1, 1, 4, 1, 2, 3, 4, 1, 1, 2},
}

func majorityElement(nums []int) int {
	var (
		count, temp int
	)
	for i := 0; i < len(nums); i++ {
		switch {
		case count == 0:
			temp = nums[i]
			count++
		case temp == nums[i]:
			count++
		default:
			count--
		}
		fmt.Println(nums[i], count, temp)
	}
	return temp
}

func main() {
	for _, v := range numbers {
		s := majorityElement(v)
		fmt.Println(s)
	}
}
   // Runtime: 12 ms, faster than 99.68% of Go online submissions for Majority Element.
   // Memory Usage: 5.9 MB, less than 83.24% of Go online submissions for Majority Element.
   ```

   

   