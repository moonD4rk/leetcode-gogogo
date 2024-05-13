## single-number

Given a **non-empty** array of integers, every element appears *twice* except for one. Find that single one.

**Note:**

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

**Example 1:**

```
Input: [2,2,1]
Output: 1
```

**Example 2:**

```
Input: [4,1,2,1,2]
Output: 4
```



### 解题方法

题目很简单，找出 array 中单独的元素，要求不使用额外空间。所以不能使用 Hash 表。

这里我们用到 `xor` 异或运算


|**输入**A|输入B|输出 A XOR B|
|:-:|:-:|:-:|
|0	|0|	0|
|0	|1|	1|
|1	|0|	1|
|1	|1|	0|

如 10 量用二进制表示为 `1010`，9用二进制表示为 `1001`，则两者取异或后的值为 `0011`  3

```go
func main() {
	var (
		a = 10
		b = 9
	)
	c := a ^ b
	fmt.Println(c)
}
```

所以答案很简单，对列表中的元素依次取异或，最后的结果为唯一单独值。

```go
func singleNumber(nums []int) int {
	var result int
	for _, v := range nums {
		result = result ^ v
	}
	return result
}
// Runtime: 4 ms, faster than 99.90% of Go online submissions for Single Number.
// Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Single Number.
```

