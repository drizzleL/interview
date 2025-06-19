package main

func replaceValueInTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodes := []*TreeNode{root}
	lvlSum := map[int]int{}
	for lvl := 1; len(nodes) != 0; lvl++ {
		var next []*TreeNode
		for _, n := range nodes {
			lvlSum[lvl] += n.Val
			if n.Left != nil {
				next = append(next, n.Left)
			}
			if n.Right != nil {
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	nodes = []*TreeNode{root}
	for lvl := 1; len(nodes) != 0; lvl++ {
		var next []*TreeNode
		for _, n := range nodes {
			var childSum int
			if n.Left != nil {
				childSum += n.Left.Val
			}
			if n.Right != nil {
				childSum += n.Right.Val
			}
			if n.Left != nil {
				n.Left.Val = lvlSum[lvl+1] - childSum
				next = append(next, n.Left)
			}
			if n.Right != nil {
				n.Right.Val = lvlSum[lvl+1] - childSum
				next = append(next, n.Right)
			}
		}
		nodes = next
	}
	root.Val = 0
	return root
}
