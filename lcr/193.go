package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val > q.Val {
		p, q = q, p
	}
	for {
		if root == p || root == q {
			return root
		}
		if root.Val > p.Val && root.Val < q.Val {
			return root
		}
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
