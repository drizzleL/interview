package main

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var idx int
	var maxVal int
	for i, num := range nums {
		if num > maxVal {
			maxVal = num
			idx = i
		}
	}
	ret := &TreeNode{
		Val:   maxVal,
		Left:  constructMaximumBinaryTree(nums[:idx]),
		Right: constructMaximumBinaryTree(nums[idx+1:]),
	}
	return ret
}
