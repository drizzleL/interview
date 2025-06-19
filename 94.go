package main

func inorderTraversal(root *TreeNode) []int {
	var ret []int
	var q []*TreeNode
	var seen []bool
	for root != nil || len(q) != 0 {
		for root != nil {
			q = append(q, root)
			seen = append(seen, false)
			root = root.Left
		}
		top := q[len(q)-1]
		q = q[:len(q)-1]
		ret = append(ret, top.Val)
		root = top.Right
	}
	return ret
}
