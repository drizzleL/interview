package main

func distributeCoins(root *TreeNode) int {
	var ret int
	var helper func(n *TreeNode) int
	helper = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		left, right := helper(n.Left), helper(n.Right)
		ret += abs(left) + abs(right)
		return n.Val + left + right - 1
	}
	helper(root)
	return ret
}
