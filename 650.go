package main

import "math"

func minSteps2(n int) int {
	if n <= 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	dp[1] = 0
	for i := 1; i < n; i++ {
		for k := 2; k*i <= n; k++ {
			dp[i*k] = min(dp[i*k], dp[i]+k)
		}
	}
	return dp[n]
}
