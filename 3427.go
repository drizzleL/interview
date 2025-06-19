package main

func subarraySum2(nums []int) int {
	presum := make([]int, len(nums)+1)
	var ret int
	for i := 0; i < len(nums); i++ {
		presum[i+1] = presum[i] + nums[i]
		start := max(0, i-nums[i])
		ret += presum[i+1] - presum[start]
	}
	return ret
}
