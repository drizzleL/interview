package main

func reverseOddLevels(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var flag bool
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		if flag {
			for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
				a, b := nodes[i], nodes[j]
				a.Val, b.Val = b.Val, a.Val
			}
		}
		var next []*TreeNode
		for _, n := range nodes {
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
		flag = !flag
	}
	return root
}
