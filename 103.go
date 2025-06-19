package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	nodes := []*TreeNode{root}
	for flag := 0; len(nodes) != 0; flag ^= 1 {
		var next []*TreeNode
		var vals []int
		for _, n := range nodes {
			vals = append(vals, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if flag == 1 {
			for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
				vals[i], vals[j] = vals[j], vals[i]
			}
		}
		ret = append(ret, vals)
		nodes = next
	}
	return ret
}
