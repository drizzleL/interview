package main

import (
	"strings"
)

func smallestFromLeaf(root *TreeNode) string {
	var helper func(n *TreeNode, prefix string) string
	helper = func(n *TreeNode, prefix string) string {
		if n == nil {
			return prefix
		}
		base := string('a' + n.Val)
		if n.Left == nil && n.Right == nil {
			return base
		}
		var s []string
		if n.Left != nil {
			s = append(s, helper(n.Left, base+prefix))
		}
		if n.Right != nil {
			s = append(s, helper(n.Right, base+prefix))
		}
		if len(s) == 1 || strings.Compare(s[0]+base+prefix, s[1]+base+prefix) < 0 {
			return s[0] + base
		}
		return s[1] + base
	}
	return helper(root, "")
}
