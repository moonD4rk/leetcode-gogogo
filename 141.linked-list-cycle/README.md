## Linked List Cycle

Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, we use an integer `pos` which represents the position (0-indexed) in the linked list where tail connects to. If `pos` is `-1`, then there is no cycle in the linked list.

**Example 1:**

```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

**示例 2：**

```
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。
```

**示例 3：**

```
输入：head = [1], pos = -1
输出：false	
解释：链表中没有环。
```

### Solution

#### 1. 将指针存储到 Map，判断是否存在

```go
func hasCycle(head *ListNode) bool {
	var temp = make(map[*ListNode]bool)
	for head != nil {
		if temp[head] {
			return true
		}
		temp[head] = true
		head = head.Next
	}
	return false
}
// Runtime: 12 ms, faster than 16.48% of Go online submissions for Linked List Cycle.
// Memory Usage: 5.2 MB, less than 5.53% of Go online submissions for Linked List Cycle.
```

