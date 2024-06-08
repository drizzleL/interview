package main

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	pre := &TreeNode{
		Left: root,
	}
	nodes := []*TreeNode{pre}
	for i := 1; i < depth; i++ {
		var next []*TreeNode
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
	for _, n := range nodes {
		left, right := n.Left, n.Right
		n.Left = &TreeNode{
			Val:  val,
			Left: left,
		}
		n.Right = &TreeNode{
			Val:   val,
			Right: right,
		}
	}
	return pre.Left
}
