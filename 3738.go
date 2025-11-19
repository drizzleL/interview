package main

func longestSubarray3(nums []int) int {
	dp := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = 1
		if i != len(nums)-1 && nums[i] <= nums[i+1] {
			dp[i] += dp[i+1]
		}
	}
	var ret int
	for i := 0; i < len(nums); i++ {
		j := dp[i] + i - 1
		ret = max(ret, j-i+1)
		if j != len(nums)-1 {
			if i == j || nums[j-1] <= nums[j+1] { // change j
				ret = max(ret, dp[i]+dp[j+1])
			}
			if j+1 < len(nums) { // change j + 1
				maxVal := nums[j]
				j += 1
				if j+1 < len(nums) && nums[j+1] >= maxVal {
					j += dp[j+1]
				}
				ret = max(ret, j-i+1)
			}
		}
	}
	return ret
}
