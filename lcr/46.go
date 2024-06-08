package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		var next []*TreeNode
		ret = append(ret, nodes[len(nodes)-1].Val)
		for _, n := range nodes {
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
