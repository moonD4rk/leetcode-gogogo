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



