package main

func wiggleMaxLength(nums []int) int {
	var helper func(i int, flag bool) int
	helper = func(i int, flag bool) int {
		if i == len(nums) {
			return 0
		}
		for i+1 < len(nums) && (nums[i+1] == nums[i] || (nums[i+1] > nums[i]) == flag) {
			i++
		}
		return 1 + helper(i+1, !flag)
	}
	return max(helper(0, false), helper(0, true))
}
