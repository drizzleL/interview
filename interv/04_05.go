package main

import "math"

func isValidBST(root *TreeNode) bool {
	var helper func(n *TreeNode, mini, maxi int) bool
	helper = func(n *TreeNode, mini, maxi int) bool {
		if n == nil {
			return true
		}
		if n.Val >= maxi || n.Val <= mini {
			return false
		}
		return helper(n.Left, mini, n.Val) && helper(n.Right, n.Val, maxi)
	}
	return helper(root, math.MinInt64, math.MaxInt64)
}
