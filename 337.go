package main

func rob(root *TreeNode) int {
	type key struct {
		Node *TreeNode
		Flag bool
	}
	cache := map[key]int{}
	var helper func(node *TreeNode, flag bool) int
	helper = func(node *TreeNode, flag bool) (ret int) {
		if node == nil {
			return 0
		}
		if c, ok := cache[key{node, flag}]; ok {
			return c
		}
		defer func() {
			cache[key{node, flag}] = ret
		}()
		ret = helper(node.Left, false) + helper(node.Right, false)
		if !flag {
			ret = max(ret, node.Val+helper(node.Left, true)+helper(node.Right, true))
		}
		return ret
	}
	return helper(root, false)
}
