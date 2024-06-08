package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil {
		return nil
	}
	dict := map[int]bool{}
	for _, d := range to_delete {
		dict[d] = true
	}
	var ret []*TreeNode
	var helper func(n *TreeNode)
	var traverse func(n *TreeNode)
	traverse = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left != nil && dict[n.Left.Val] {
			helper(n.Left.Left)
			helper(n.Left.Right)
			n.Left = nil
		} else {
			traverse(n.Left)
		}
		if n.Right != nil && dict[n.Right.Val] {
			helper(n.Right.Left)
			helper(n.Right.Right)
			n.Right = nil
		} else {
			traverse(n.Right)
		}
	}
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}
		if dict[n.Val] {
			helper(n.Left)
			helper(n.Right)
			return
		}
		ret = append(ret, n)
		traverse(n)
	}
	helper(root)
	return ret
}
