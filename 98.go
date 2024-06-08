package main

import "math"

func isValidBST(root *TreeNode) bool {
	var helper func(node *TreeNode, maxVal, minVal int) bool
	helper = func(node *TreeNode, maxVal, minVal int) bool {
		if node == nil {
			return true
		}
		if node.Val >= maxVal || node.Val <= minVal {
			return false
		}
		return helper(node.Left, node.Val, minVal) && helper(node.Right, maxVal, node.Val)
	}
	return helper(root, math.MaxInt32+1, math.MinInt32-1)
}
