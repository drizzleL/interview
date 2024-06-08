package main

import "math"

func recoverTree(root *TreeNode) {
	var first, second *TreeNode
	pre := &TreeNode{
		Val: math.MinInt64,
	}
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		if first == nil && pre.Val >= node.Val {
			first = pre
		}
		if first != nil && pre.Val >= node.Val {
			second = node
		}
		pre = node
		helper(node.Right)
	}
	helper(root)
	first.Val, second.Val = second.Val, first.Val
}
