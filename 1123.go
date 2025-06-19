package main

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var next []*TreeNode
		for _, node := range nodes {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		if len(next) == 0 {
			break
		}
		nodes = next
	}
	first, last := nodes[0], nodes[len(nodes)-1]
	return getLca(root, first, last)
}

func getLca(root *TreeNode, a, b *TreeNode) *TreeNode {
	var ret *TreeNode
	var find func(node *TreeNode, a, b *TreeNode) (flag1, flag2 bool)
	find = func(node *TreeNode, a, b *TreeNode) (flag1, flag2 bool) {
		if node == nil {
			return
		}
		if ret != nil {
			return
		}
		if node == a {
			flag1 = true
		}
		if node == b {
			flag2 = true
		}
		left1, left2 := find(node.Left, a, b)
		right1, right2 := find(node.Right, a, b)
		flag1 = flag1 || left1 || right1
		flag2 = flag2 || left2 || right2
		if ret == nil && flag1 && flag2 {
			ret = node
		}
		return
	}
	find(root, a, b)
	return ret
}
