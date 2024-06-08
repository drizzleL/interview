package main

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	var ret int
	for len(nodes) != 0 {
		var next []*TreeNode
		ret = 0
		for _, n := range nodes {
			ret += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	return ret
}
