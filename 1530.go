package main

func countPairs(root *TreeNode, distance int) int {
	if root == nil {
		return 0
	}
	var ret int
	var helper func(node *TreeNode) []int
	helper = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{1}
		}
		left, right := helper(node.Left), helper(node.Right)
		for i := 0; i < len(left); i++ {
			for j := 0; j < len(right) && i+j+2 <= distance; j++ {
				ret += left[i] * right[j]
			}
		}
		ans := make([]int, min(distance-1, max(len(left), len(right))+1))
		for i := 0; i < len(left) && i < distance-1 && i+1 < len(ans); i++ {
			ans[i+1] += left[i]
		}
		for i := 0; i < len(right) && i < distance-1 && i+1 < len(ans); i++ {
			ans[i+1] += right[i]
		}
		return ans
	}
	helper(root)
	return ret
}
