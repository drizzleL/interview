package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == q || root == p {
		return root
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor2(root.Left, p, q)
	}
	return lowestCommonAncestor2(root.Right, p, q)
}
