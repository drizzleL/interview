package main

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p == nil || root == nil {
		return nil
	}
	if p.Right != nil { // next in p.Right
		node := p.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	}
	var lastNode *TreeNode
	node := root
	for node != nil && node.Val != p.Val {
		if node.Val < p.Val {
			node = node.Right
		} else if node.Val > p.Val {
			lastNode = node
			node = node.Left
		} else {
			break
		}
	}
	return lastNode
}
