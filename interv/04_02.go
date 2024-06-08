package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	midIdx := (len(nums) - 1) / 2
	node := &TreeNode{Val: nums[midIdx]}
	node.Left = sortedArrayToBST(nums[:midIdx])
	node.Right = sortedArrayToBST(nums[midIdx+1:])
	return node
}
