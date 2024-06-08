package main

func sumNumbers(root *TreeNode) int {
	var ret int
	var helper func(node *TreeNode, now int)
	helper = func(node *TreeNode, now int) {
		if node == nil {
			return
		}
		now = now*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ret += now
			return
		}
		helper(node.Left, now)
		helper(node.Right, now)
	}
	helper(root, 0)
	return ret
}
