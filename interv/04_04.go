package main

func isBalanced(root *TreeNode) bool {
	var getHeight func(n *TreeNode) int
	getHeight = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		h1, h2 := getHeight(n.Left), getHeight(n.Right)
		if h1 > h2 {
			h1, h2 = h2, h1
		}
		if h1 == -1 || h2 == -1 {
			return -1
		}
		if h2-h1 > 1 {
			return -1
		}
		return h2 + 1
	}
	return getHeight(root) != -1
}
