package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	// 初始化一个slow链表，将原始节点挂载到此节点之后
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

func NewNode() *ListNode {
	return &ListNode{
		Val:  0,
		Next: nil,
	}
}

func (l *ListNode) Append(num int) {
	newNode := &ListNode{
		Val:  num,
		Next: nil,
	}
	for iter := l; iter != nil; iter = iter.Next {
		if iter.Next == nil {
			iter.Next = newNode
			break
		}
	}
}

func (l *ListNode) Range() {
	for iter := l; iter != nil; iter = iter.Next {
		fmt.Printf("%v -> ", iter.Val)
	}
	fmt.Println()
	fmt.Printf("-----range is done\n")
}

func main() {
	l := NewNode()
	l.Append(1)
	l.Append(2)
	l.Append(6)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.Range()
	m := removeElements(l, 6)
	m.Range()
}
