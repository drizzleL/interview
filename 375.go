package main

import (
	"math"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
		for j := i + 1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	for gap := 1; gap < n; gap++ {
		for i := 1; i+gap <= n; i++ {
			for j := i; j <= i+gap; j++ {
				dp[i][i+gap] = min(dp[i][i+gap], j+max(dp[i][j-1], dp[j+1][i+gap]))
			}
		}
	}
	return dp[1][n]
}
