package main

import "math"

func paintWalls(cost []int, time []int) int {
	dp := make([]int, len(cost)+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := range cost {
		c, t := cost[i], time[i]
		for j := len(dp) - 1; j > 0; j-- {
			dp[j] = min(dp[j], dp[max(0, j-t-1)]+c)
		}
	}
	return dp[len(cost)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
