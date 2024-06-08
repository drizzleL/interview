package main

func sumOfLeftLeaves(root *TreeNode) int {
	var helper func(node *TreeNode, isLeft bool) int
	helper = func(node *TreeNode, isLeft bool) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			if isLeft {
				return node.Val
			}
			return 0
		}
		return helper(node.Left, true) + helper(node.Right, false)
	}
	return helper(root, false)
}
