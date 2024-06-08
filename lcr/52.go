package main

func increasingBST(root *TreeNode) *TreeNode {
	var helper func(node *TreeNode) (*TreeNode, *TreeNode)
	helper = func(node *TreeNode) (*TreeNode, *TreeNode) {
		if node.Left == nil && node.Right == nil {
			return node, node
		}
		retRoot, retTail := node, node
		if node.Left != nil {
			lRoot, lTail := helper(node.Left)
			retRoot = lRoot
			lTail.Right = node
		}
		if node.Right != nil {
			rRoot, rTail := helper(node.Right)
			node.Right = rRoot
			retTail = rTail
		}
		node.Left = nil
		return retRoot, retTail
	}
	ret, _ := helper(root)
	return ret
}
