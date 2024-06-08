package main

func isBalanced(root *TreeNode) bool {
	var flag bool
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lh, rh := helper(node.Left), helper(node.Right)
		if lh > rh {
			lh, rh = rh, lh
		}
		if rh-lh > 1 {
			flag = true
			return 0
		}
		return rh + 1
	}
	helper(root)
	return !flag
}
