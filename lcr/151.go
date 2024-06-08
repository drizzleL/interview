package main

func decorateRecord2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ret [][]int
	var flag bool
	for len(nodes) != 0 {
		var lvl []int
		var next []*TreeNode
		for _, n := range nodes {
			lvl = append(lvl, n.Val)
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if flag {
			for i, j := 0, len(lvl); i < j; i, j = i+1, j-1 {
				lvl[i], lvl[j] = lvl[j], lvl[i]
			}
		}
		flag = !flag
		nodes = next
		ret = append(ret, lvl)
	}
	return ret
}
