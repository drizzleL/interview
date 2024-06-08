package main

func calculateDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(calculateDepth(root.Left), calculateDepth(root.Right))
}
