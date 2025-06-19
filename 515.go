package main

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		var next []*TreeNode
		tmp := nodes[0].Val
		for _, node := range nodes {
			tmp = max(tmp, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		ret = append(ret, tmp)
		nodes = next
	}
	return ret
}
