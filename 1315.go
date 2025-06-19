package main

func sumEvenGrandparent(root *TreeNode) int {
	var ret int
	var helper func(node *TreeNode, evenP, evenGp bool)
	helper = func(node *TreeNode, evenP, evenGp bool) {
		if node == nil {
			return
		}
		if evenGp {
			ret += node.Val
		}
		helper(node.Left, node.Val%2 == 0, evenP)
		helper(node.Right, node.Val%2 == 0, evenP)
	}
	helper(root, false, false)
	return ret
}
