## Same Tree

Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

**Example 1:**

```
Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
```

**Example 2:**

```
Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
```

**Example 3:**

```
Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
```

### 解决办法

依次判断节点是不是nil，节点不为 nil 情况下判断节点的值是否相等

```go
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
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Same Tree.
// Memory Usage: 2.4 MB, less than 5.02% of Go online submissions for Same Tree.
```

