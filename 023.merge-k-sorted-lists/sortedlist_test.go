package _23_merge_k_sorted_lists

import (
	"testing"
)

var (
	l1 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	l2 = &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l3 = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
)

func TestMergeKLists(t *testing.T) {
	sorted := MergeKLists([]*ListNode{l1, l2, l3})
	if sorted == nil {

	}
}
