package main

func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for _, num := range nums {
		for i := target; i >= num; i-- {
			if dp[i-num] == -1 {
				continue
			}
			dp[i] = max(dp[i], dp[i-num]+1)
		}
	}
	return dp[target]
}
