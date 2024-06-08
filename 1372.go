package main

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	var helper func(n *TreeNode, curr int, dir int)
	helper = func(n *TreeNode, curr int, dir int) {
		if n == nil {
			return
		}
		ret = max(ret, curr+1)
		if dir == 0 {
			helper(n.Right, curr+1, 1)
			helper(n.Left, 0, 0)
		} else {
			helper(n.Left, curr+1, 0)
			helper(n.Right, 0, 1)
		}

	}
	helper(root.Left, 0, 0)
	helper(root.Right, 0, 1)
	return ret
}
