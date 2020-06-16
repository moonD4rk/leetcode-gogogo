# 两数相加

[题目链接](https://leetcode.com/problems/add-two-numbers/)

## 题目
给定两个链表，按照正常的加法法则返回他们相加后的结果。

```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807
```

答案

```go

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	node3 := l3
	a := 0
	for (l1 != nil) || (l2 != nil) || a > 0 {
		node3.Next = new(ListNode)
		node3 = node3.Next
		b := 0
		c := 0
		if l1 != nil {
			b = l1.Val
		}
		if l2 != nil {
			c = l2.Val
		}
		node3.Val = (a + b + c) % 10
		a = (a + b + c) / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return l3.Next
}
```



对于数据结构的处理还是有一些不清除的地方，需要继续学习。



