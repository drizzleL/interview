package main

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret, sum int
	nodes := []*TreeNode{root}
	for i := 1; len(nodes) != 0; i++ {
		var next []*TreeNode
		var lvlSum int
		for _, n := range nodes {
			lvlSum += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		if lvlSum > sum {
			ret = i
		}
		nodes = next
	}
	return ret
}
