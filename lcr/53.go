package main

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		node := p.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	node := root
	var ret *TreeNode
	for node != nil {
		if node.Val > p.Val {
			ret = node
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return ret
}
