package main

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := root.Val
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := max(helper(node.Left), 0), max(helper(node.Right), 0)
		val := l + r + node.Val
		if val > ret {
			ret = val
		}
		return max(l, r) + node.Val
	}
	helper(root)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
