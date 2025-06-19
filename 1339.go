package main

func maxProduct5(root *TreeNode) int {
	var sum int
	var ret int
	var getSum func(node *TreeNode) int
	getSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		val := node.Val + getSum(node.Left) + getSum(node.Right)
		ret = max(ret, val*(sum-val))
		return val
	}
	sum = getSum(root)
	getSum(root)
	ret %= 1e9 + 7
	return ret
}
