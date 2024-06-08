package main

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var next []*TreeNode
		maxVal := nodes[0].Val
		for _, n := range nodes {
			if n.Val > maxVal {
				maxVal = n.Val
			}
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
		ret = append(ret, maxVal)
	}
	return ret
}
