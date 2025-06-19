package main

func countPartitions(nums []int, k int) int {
	dp := make([]int, len(nums))
	presum := make([]int, len(nums)+1)
	presum[0] = 1
	var mins, maxs []int
	check := func(val int) bool {
		if len(maxs) != 0 && maxs[0]-val > k {
			return false
		}
		if len(mins) != 0 && val-mins[0] > k {
			return false
		}
		return true
	}
	var head int
	for i := 0; i < len(nums); i++ {
		for !check(nums[i]) {
			if len(mins) != 0 && mins[0] == nums[head] {
				mins = mins[1:]
			}
			if len(maxs) != 0 && maxs[0] == nums[head] {
				maxs = maxs[1:]
			}
			head++
		}
		for len(maxs) != 0 && maxs[len(maxs)-1] < nums[i] {
			maxs = maxs[:len(maxs)-1]
		}
		maxs = append(maxs, nums[i])
		for len(mins) != 0 && mins[len(mins)-1] > nums[i] {
			mins = mins[:len(mins)-1]
		}
		mins = append(mins, nums[i])
		dp[i] = presum[i]
		if head != 0 {
			dp[i] -= presum[head-1]
		}
		if dp[i] < 0 {
			dp[i] += 1e9 + 7
		}
		dp[i] %= 1e9 + 7
		presum[i+1] = presum[i] + dp[i]
		presum[i+1] %= 1e9 + 7
	}
	return dp[len(dp)-1]
}
