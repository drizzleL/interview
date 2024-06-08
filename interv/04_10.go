package main

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil {
		return t2 == nil
	}
	var check func(a, b *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}
		if a.Val != b.Val {
			return false
		}
		return check(a.Left, b.Left) && check(a.Right, b.Right)
	}
	q := []*TreeNode{t1}
	for len(q) != 0 {
		top := q[len(q)-1]
		q = q[:len(q)-1]
		if check(top, t2) {
			return true
		}
		if top.Left != nil {
			q = append(q, top.Left)
		}
		if top.Right != nil {
			q = append(q, top.Right)
		}
	}
	return false
}
