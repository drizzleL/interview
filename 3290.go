package main

import "math"

func maxScore9(a []int, b []int) int64 {
	dp := make([][]int, len(b))
	for i := range dp {
		dp[i] = make([]int, 4)
		for j := range dp[i] {
			dp[i][j] = math.MinInt64
		}
	}
	for i := 0; i < len(b); i++ {
		dp[i][0] = a[0] * b[i]
		if i > 0 {
			dp[i][0] = max(dp[i][0], dp[i-1][0])
		}
	}
	for j := 1; j < 4; j++ {
		for i := j; i < len(b); i++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+a[j]*b[i])
		}
	}
	return int64(dp[len(b)-1][3])
}
