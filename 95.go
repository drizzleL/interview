package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	var helper func(start, end int) []*TreeNode
	helper = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var ret []*TreeNode
		for mid := start; mid <= end; mid++ {
			for _, lChild := range helper(start, mid-1) {
				for _, rChild := range helper(mid+1, end) {
					node := &TreeNode{Val: mid}
					node.Left = lChild
					node.Right = rChild
					ret = append(ret, node)
				}
			}
		}
		return ret
	}
	return helper(1, n)
}
