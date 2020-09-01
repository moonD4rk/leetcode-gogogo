package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	switch {
	case p == nil && q == nil:
		return true
	case p != nil && q == nil:
		return false
	case p == nil && q != nil:
		return false
	case p != nil && q != nil:
		if p.Val != q.Val {
			return false
		} else {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	}
	return false
}

func main() {
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	t2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	b := isSameTree(t1, t2)
	fmt.Println(b)
}
