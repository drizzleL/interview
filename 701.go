package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	v := &TreeNode{Val: val}
	if root == nil {
		return v
	}
	check := func(n *TreeNode) (*TreeNode, *TreeNode) {
		if n == nil {
			return v, nil
		}
		return n, n
	}
	for node := root; node != nil; {
		if val < node.Val {
			node.Left, node = check(node.Left)
		} else {
			node.Right, node = check(node.Right)
		}
	}
	return root
}
