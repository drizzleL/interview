package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	var helper func(a, b *TreeNode) bool
	helper = func(a, b *TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		return helper(a.Left, b.Left) && helper(a.Right, b.Right)
	}
	nodes := []*TreeNode{A}
	for len(nodes) > 0 {
		var newnodes []*TreeNode
		for _, n := range nodes {
			if helper(n, B) {
				return true
			}
			if n.Left != nil {
				newnodes = append(newnodes, n.Left)
			}
			if n.Right != nil {
				newnodes = append(newnodes, n.Right)
			}
		}
		nodes = newnodes
	}
	return false
}
