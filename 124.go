package main

import "math"

func maxPathSum(root *TreeNode) int {
	ret := math.MinInt32
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ret = max(ret, node.Val)
		left, right := helper(node.Left), helper(node.Right)
		if left < 0 && right < 0 {
			return node.Val
		}
		val := node.Val
		if left > 0 {
			val += left
		}
		if right > 0 {
			val += right
		}
		ret = max(ret, val)
		maxLine := node.Val + max(left, right)
		return maxLine
	}
	helper(root)
	return ret
}
