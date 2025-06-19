package main

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) []string {
	var ret []string
	var helper func(node *TreeNode, pre []string)
	helper = func(node *TreeNode, pre []string) {
		if node == nil {
			return
		}
		pre = append(pre, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			ret = append(ret, strings.Join(pre, "->"))
			return
		}
		helper(node.Left, pre)
		helper(node.Right, pre)
		return
	}
	helper(root, nil)
	return ret
}
