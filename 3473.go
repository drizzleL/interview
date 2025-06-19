package main

import (
	"math"
)

func maxSum3(nums []int, k int, m int) int {
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 0
		for j := 1; j <= k; j++ {
			dp[i][j] = math.MinInt32
		}
	}
	presum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	for j := 1; j <= k; j++ {
		best := math.MinInt32
		for i := j * m; i <= len(nums); i++ {
			start := i - m
			if start >= 0 {
				best = max(best, dp[start][j-1]-presum[start])
			}
			if i == j*m {
				dp[i][j] = presum[i] + best
			} else {
				dp[i][j] = max(dp[i-1][j], presum[i]+best)
			}
		}
	}
	return dp[len(nums)][k]
}
