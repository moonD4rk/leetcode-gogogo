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

func main() {

}
