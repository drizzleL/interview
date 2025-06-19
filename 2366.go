package main

func minimumReplacement(nums []int) int64 {
	var ret int
	maxVal := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= maxVal {
			maxVal = nums[i]
			continue
		}
		chunks := (nums[i] + maxVal - 1) / maxVal
		ret += chunks - 1
		maxVal = nums[i] / chunks
	}
	return int64(ret)
}
