package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	out := l3
	log.Printf("%p %p\n", l3, out)
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
	log.Printf("%p %p\n", l3, out)
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
	log.Println("l1")
	log.Println()
	l3 := mergeTwoLists(l1, l2)
	log.Println()
	l3.rangeList()
}

func (l *ListNode) rangeList() {
	for l != nil {
		log.Print(l.Val, "> ")
		l = l.Next
	}
	log.Println()
}
