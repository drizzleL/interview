package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var helper func(node *TreeNode) (*TreeNode, bool, bool)
	helper = func(node *TreeNode) (lca *TreeNode, left, right bool) {
		if node == nil {
			return nil, false, false
		}
		if node == p {
			left = true
		}
		if node == q {
			right = true
		}
		lchildLca, lchildLeft, lchildRight := helper(node.Left)
		if lchildLca != nil { // early break
			return lchildLca, false, false
		}
		rchildLca, rchildLeft, rchildRight := helper(node.Right)
		if rchildLca != nil { // early break
			return rchildLca, false, false
		}
		left = left || lchildLeft || rchildLeft
		right = right || lchildRight || rchildRight
		if left && right {
			lca = node
		}
		return
	}
	r, _, _ := helper(root)
	return r
}
