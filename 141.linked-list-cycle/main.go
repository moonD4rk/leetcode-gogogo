package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func hasCycle2Point(node *ListNode) bool {
	fast := node
	slow := node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {

}
