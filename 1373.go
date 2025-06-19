package main

import "math"

func maxSumBST(root *TreeNode) int {
	var ret int
	var dfs func(x *TreeNode) (sum, minVal, maxVal int)
	dfs = func(x *TreeNode) (sum, minVal, maxVal int) {
		if x == nil {
			return 0, math.MinInt32, math.MaxInt32
		}
		sum, minVal, maxVal = x.Val, x.Val, x.Val
		leftSum, leftMin, leftMax := dfs(x.Left)
		rightSum, rightMin, rightMax := dfs(x.Right)
		if leftSum == math.MinInt32 || rightSum == math.MinInt32 {
			return math.MinInt32, math.MinInt32, math.MaxInt32
		}
		if leftMax != math.MaxInt32 && x.Val <= leftMax || rightMin != math.MinInt32 && x.Val >= rightMin {
			return math.MinInt32, 0, 0
		}
		sum += leftSum + rightSum
		if leftMin != math.MinInt32 {
			minVal = leftMin
		}
		if rightMax != math.MaxInt32 {
			maxVal = rightMax
		}
		ret = max(ret, sum)
		return
	}
	dfs(root)
	return ret
}
