package main

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var helper func(node *TreeNode, dict map[int]int)
	var ret int
	helper = func(node *TreeNode, dict map[int]int) {
		newdict := map[int]int{}
		for k, v := range dict {
			newdict[k+node.Val] += v
		}
		newdict[node.Val] += 1
		ret += newdict[targetSum]
		if node.Left != nil {
			helper(node.Left, newdict)
		}
		if node.Right != nil {
			helper(node.Right, newdict)
		}
	}
	helper(root, nil)
	return ret
}

func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var helper func(node *TreeNode, sum int)
	var ret int
	dict := map[int]int{
		0: 1,
	}
	helper = func(node *TreeNode, sum int) {
		currsum := sum + node.Val
		dict[currsum]++
		ret += dict[targetSum-currsum]
		if node.Left != nil {
			helper(node.Left, currsum)
		}
		if node.Right != nil {
			helper(node.Right, currsum)
		}
		dict[currsum]--
	}
	helper(root, 0)
	return ret
}
