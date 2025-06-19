package main

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}
	var xParent, yParent *TreeNode
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var next []*TreeNode
		for _, n := range nodes {
			if n.Left != nil {
				if n.Left.Val == x {
					xParent = n
				}
				if n.Left.Val == y {
					yParent = n
				}
				next = append(next, n.Left)
			}
			if n.Right != nil {
				if n.Right.Val == x {
					xParent = n
				}
				if n.Right.Val == y {
					yParent = n
				}
				next = append(next, n.Right)
			}
		}
		if xParent != nil && yParent != nil {
			return xParent != yParent
		}
		if xParent != nil || yParent != nil {
			return false
		}
		nodes = next
	}
	return false
}
