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

// Definition for singly-linked list.
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
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 2.6 MB, less than 76.41% of Go online submissions for Reverse Linked List.

func main() {
	m := reverseList(l)
	for m != nil {
		fmt.Println(m)
		m = m.Next
	}
}
