package main

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	dict := map[string]int{}
	var ret []*TreeNode
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		key := strconv.Itoa(node.Val) + "," + dfs(node.Left) + "," + dfs(node.Right)
		dict[key] += 1
		if dict[key] == 2 {
			ret = append(ret, root)
		}
		return key
	}
	dfs(root)
	return ret
}
