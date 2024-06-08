package main

func pathSum2(root *TreeNode, targetSum int) int {
	var presum int
	var stack []*TreeNode
	var used []bool
	dict := map[int]int{
		0: 1,
	}
	var ret int
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			used = append(used, false)
			presum += root.Val
			ret += dict[presum-targetSum]
			dict[presum] += 1
			root = root.Left
		}
		top := stack[len(stack)-1]
		if top.Right != nil && !used[len(used)-1] {
			used[len(used)-1] = true
			root = top.Right
			continue
		}
		stack = stack[:len(stack)-1]
		used = used[:len(used)-1]
		dict[presum] -= 1
		presum -= top.Val
	}
	return ret
}
