package main

func decorateRecord(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	var ret []int
	for len(nodes) != 0 {
		var newnodes []*TreeNode
		for _, n := range nodes {
			ret = append(ret, n.Val)
			if n.Left != nil {
				newnodes = append(newnodes, n.Left)
			}
			if n.Right != nil {
				newnodes = append(newnodes, n.Right)
			}
		}
		nodes = newnodes
	}
	return ret
}
