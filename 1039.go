package main

import "math"

func minScoreTriangulation(values []int) int {
	dp := make([][]int, len(values))
	for i := range dp {
		dp[i] = make([]int, len(values))
	}
	for d := 2; d < len(values); d++ {
		for i := 0; i+d < len(values); i++ {
			j := i + d
			dp[i][j] = math.MaxInt32
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+values[i]*values[j]*values[k])
			}
		}
	}
	return dp[0][len(values)-1]
}
