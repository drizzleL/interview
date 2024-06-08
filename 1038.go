package main

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var helper func(n *TreeNode, presum int) int
	helper = func(n *TreeNode, presum int) int {
		if n == nil {
			return 0
		}
		var sum int
		sum += n.Val
		sum += helper(n.Right, presum)
		n.Val = presum + sum
		sum += helper(n.Left, sum+presum)
		return sum
	}
	helper(root, 0)
	return root
}
