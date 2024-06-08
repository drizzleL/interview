package main

func balanceBST(root *TreeNode) *TreeNode {
	var nodeToVals func(n *TreeNode) []int
	nodeToVals = func(n *TreeNode) []int {
		if n == nil {
			return nil
		}
		var ret []int
		ret = append(ret, nodeToVals(n.Left)...)
		ret = append(ret, n.Val)
		ret = append(ret, nodeToVals(n.Right)...)
		return ret
	}
	var valsToTree func(vals []int) *TreeNode
	valsToTree = func(vals []int) *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		idx := len(vals) / 2
		node := &TreeNode{
			Val: vals[idx],
		}
		node.Left = valsToTree(vals[:idx])
		if idx+1 <= len(vals) {
			node.Right = valsToTree(vals[idx+1:])
		}
		return node
	}
	return valsToTree(nodeToVals(root))
}
