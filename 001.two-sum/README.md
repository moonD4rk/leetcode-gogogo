# 两数相加

## 题目

[题目链接](https://leetcode.com/problems/two-sum/)

给你一个全是数字的列表，当列表中两个数之和等于一个特定数字时，返回这两个数的索引。

你可以假设每个输入都只有一个准确的结果，而且你不会使用两次相同的元素。

例如

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解决方法

`A + B = target`  等同于 `B = target - A`，`target `已经给出的情况下，遍历数组中的每个元素，判断 `target` 减去当前元素的结果在不在列表中可解。而判断元素一个值在不在一个列表中的最佳方案是使用`Hashmap`，查询、插入的时间复杂度为 O(1)。因此最终代码如下：

```go
func twoSum(nums []int, target int) []int {
	var tempMap = make(map[int]int, len(nums))
	//for i, v := range nums {
	//	tempMap[target-v] = i
	//}
	for i, v := range nums {
		if _, ok := tempMap[v]; ok {
			return []int{tempMap[v], i}
		}
		tempMap[target-v] = i
	}
	return nil
}
```

