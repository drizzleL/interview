package main

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	var helper func(node *TreeNode, preVal int)
	helper = func(node *TreeNode, preVal int) {
		if node == nil {
			return
		}
		if node.Val >= preVal {
			ret += 1
		}
		preVal = max(preVal, node.Val)
		helper(node.Left, preVal)
		helper(node.Right, preVal)
	}
	helper(root, root.Val)
	return ret
}
