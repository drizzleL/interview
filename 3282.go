package main

func findMaximumScore(nums []int) int64 {
	var idx, ret int
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[idx] {
			continue
		}
		ret += nums[idx] * (i - idx)
		idx = i
	}
	ret += nums[idx] * (len(nums) - 1 - idx)
	return int64(ret)
}
