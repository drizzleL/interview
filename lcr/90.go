package main

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}
