### Remove Linked List Elements

Remove all elements from a linked list of integers that have value ***val\***.

**Example:**

```
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
```



### 解题

```go
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// 初始化一个节点，将原始节点挂载到此节点之后
	res := &ListNode{0, head}
	slow := res
	node := head
	// 遍历 node 链表
	for node != nil {
		// 如果 node 的值和要删除的值相同
		if node.Val == val {
			// 则通过指向删除此节点
			slow.Next = node.Next
		} else {
			// 不相等，将新slow向后移一位
			slow = node
		}
		node = node.Next
	}
	return res.Next
}
```

