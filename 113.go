package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val != targetSum {
			return nil
		}
		return [][]int{{root.Val}}
	}
	var ret [][]int
	if root.Left != nil {
		for _, v := range pathSum(root.Left, targetSum-root.Val) {
			ret = append(ret, append([]int{root.Val}, v...))
		}
	}
	if root.Right != nil {
		for _, v := range pathSum(root.Right, targetSum-root.Val) {
			ret = append(ret, append([]int{root.Val}, v...))
		}
	}
	return ret
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
