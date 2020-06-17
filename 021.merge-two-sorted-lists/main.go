package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 = new(ListNode)
	var out = l3
	fmt.Printf("%p %p\n", l3, out)
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
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

	if l1 != nil {
		l3.Next = l1
	} else if l2 != nil {
		l3.Next = l2
	}
	fmt.Printf("%p %p\n", l3, out)
	return out.Next
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Val = 1
	l1.Next = &ListNode{
		2,
		nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2.Val = 1
	l2.Next = &ListNode{
		7,
		nil,
	}
	l2.Next.Next = &ListNode{
		Val:  8,
		Next: nil,
	}
	fmt.Println("l1")
	fmt.Println()
	l3 := mergeTwoLists(l1, l2)
	fmt.Println()
	l3.rangeList()
}

func (l *ListNode) rangeList() {
	for l != nil {
		fmt.Print(l.Val, "> ")
		l = l.Next
	}
	fmt.Println()
}
