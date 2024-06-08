package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ret *TreeNode
	var helper func(n *TreeNode) (bool, bool)
	helper = func(n *TreeNode) (bool, bool) {
		if n == nil || ret != nil {
			return false, false
		}
		found1, found2 := n == p, n == q
		lf1, lf2 := helper(n.Left)
		rf1, rf2 := helper(n.Right)
		if ret != nil {
			return false, false
		}
		found1 = found1 || lf1 || rf1
		found2 = found2 || lf2 || rf2
		if found1 && found2 {
			ret = n
		}
		return found1, found2
	}
	helper(root)
	return ret
}
