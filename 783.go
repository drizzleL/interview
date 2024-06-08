package main

import "math"

func minDiffInBST(root *TreeNode) int {
	ret := math.MaxInt32
	var helper func(l, r int, n *TreeNode)
	helper = func(l, r int, n *TreeNode) {
		if n == nil {
			return
		}
		ret = min(ret, r-n.Val)
		ret = min(ret, n.Val-l)
		helper(l, n.Val, n.Left)
		helper(n.Val, r, n.Right)
	}
	helper(math.MinInt32, math.MaxInt32, root)
	return ret
}
