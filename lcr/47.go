package main

func pruneTree(root *TreeNode) *TreeNode {
	var helper func(node *TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left = helper(node.Left)
		node.Right = helper(node.Right)
		if node.Left == nil && node.Right == nil && node.Val == 0 {
			return nil
		}
		return node
	}
	return helper(root)
}
