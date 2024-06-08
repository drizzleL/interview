package main

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var helper func(n *TreeNode, greater int) int
	helper = func(n *TreeNode, greater int) int {
		if n == nil {
			return 0
		}
		sum := n.Val
		sum += helper(n.Right, greater)
		n.Val = sum + greater
		sum += helper(n.Left, sum+greater)
		return sum
	}
	helper(root, 0)
	return root
}
