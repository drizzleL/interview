package main

func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	nodes := []*TreeNode{root}
	flag := 1
	check := func(n *TreeNode, lastVal int) bool {
		if n.Val%2 != flag {
			return false
		}
		if lastVal != 0 {
			if flag == 1 && n.Val <= lastVal {
				return false
			}
			if flag == 0 && n.Val >= lastVal {
				return false
			}
		}
		return true
	}
	for len(nodes) != 0 {
		var next []*TreeNode
		var lastVal int
		for _, n := range nodes {
			if !check(n, lastVal) {
				return false
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
			lastVal = n.Val
		}
		nodes = next
		flag = 1 - flag
	}
	return true
}
