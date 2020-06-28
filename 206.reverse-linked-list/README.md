## Reverse Linked List

Reverse a singly linked list.

**Example:**

```
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

**Follow up:**

A linked list can be reversed either iteratively or recursively. Could you implement both?

### 解题方法

#### 1. 迭代

迭代的核心思想是，要将新链表的值赋给当前链表的下一个值 `list.next = prefix` ， 这样 `prefix` 每次都会迭代变化如下，这样就实现了当前节点的下一个节点 `list.next`指向了上一个节点 `prefix`

```
list: 1->2->3->4->5->nil
list.next = prefix  [1->nil]
list.next = prefix  [2->1->nil]
list.next = prefix  [3->2->1->nil]
```

为了让迭代能进行下去，需要一个临时链表 `suffix` 用来存放 list 的值，

`suffix = list.next.` ...  `list = suffix`。

注意初始化 prefix 时不要使用 `new()`函数，使用 `new` 初始化出 `0`值

```go
package main

import "fmt"

var l = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	},
}

type ListNode struct {
	Val  int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	var prefix *ListNode
	suffix := new(ListNode)
	for head != nil {
		suffix = head.Next
		head.Next = prefix
		prefix = head
		head = suffix
	}
	return prefix
}

func main() {
	m := reverseList(l)
	for m != nil {
		fmt.Println(m)
		m = m.Next
	}
}
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 76.41% of Go online submissions for Reverse Linked List.


```


#### 2. 递归


