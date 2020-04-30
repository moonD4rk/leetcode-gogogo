package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {
	l3 := new(ListNode)
	fmt.Printf("%p\n", l3)
	node3 := l3
	fmt.Printf("%p\n", node3)
	node3.Next = new(ListNode)
	node3.Next = nil
	fmt.Printf("%p\n", l3.Next)
	fmt.Printf("%p\n", node3.Next)
}
