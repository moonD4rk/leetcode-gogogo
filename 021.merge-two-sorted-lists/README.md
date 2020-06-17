# Merge Two Sorted Lists

Merge two sorted linked lists and return it as a new sorted list. The new list should be made by splicing together the nodes of the first two lists.

Example:
```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

```



### 解决办法

遍历两个链表，比较节点大小，节点大的每次往后走一步，直到其中一个链表遍历完，将另一个链表的后续的所有值追加到结果上。

```go
var l3 = new(ListNode)
for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l3.Next = l1
			l1 = l1.Next
			l3 = l3.Next
		} else {
			l3.Next = l2
			l2 = l2.Next
			l3 = l3.Next
		}
		l3.rangeList()
	}
```

