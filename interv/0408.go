package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var ret *TreeNode
	var helper func(*TreeNode) (bool, bool)
	helper = func(node *TreeNode) (bool, bool) {
		if node == nil {
			return false, false
		}
		if ret != nil {
			return false, false
		}
		ok1, ok2 := node == p, node == q
		lok1, lok2 := helper(node.Left)
		rok1, rok2 := helper(node.Right)
		ok1 = ok1 || lok1 || rok1
		ok2 = ok2 || lok2 || rok2
		if ok1 && ok2 && ret == nil {
			ret = node
			return false, false
		}
		return ok1, ok2
	}
	helper(root)
	return ret
}
