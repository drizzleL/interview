package main

func pathSum(root *TreeNode, sum int) int {
	var ret int
	dict := map[int]int{
		0: 1,
	}
	var helper func(node *TreeNode, tmp int)
	helper = func(node *TreeNode, tmp int) {
		if node == nil {
			return
		}
		tmp += node.Val
		ret += dict[tmp-sum]
		dict[tmp]++
		helper(node.Left, tmp)
		helper(node.Right, tmp)
		dict[tmp]--
	}
	helper(root, 0)
	return ret
}
