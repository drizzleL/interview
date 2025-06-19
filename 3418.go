package main

import (
	"math"
)

func maximumAmount(coins [][]int) int {
	m, n := len(coins), len(coins[0])
	dp := make([][][3]int, m)
	for i := range dp {
		dp[i] = make([][3]int, n)
		for j := range dp[i] {
			dp[i][j] = [3]int{math.MinInt32, math.MinInt32, math.MinInt32}
		}
	}
	dp[0][0] = [3]int{coins[0][0], coins[0][0], coins[0][0]}
	if coins[0][0] < 0 {
		dp[0][0][1] = 0
		dp[0][0][2] = 0
	}
	helper := func(curr [3]int, last [3]int, c int) [3]int {
		curr[0] = max(curr[0], last[0]+c)
		curr[1] = max(curr[1], last[1]+c)
		curr[2] = max(curr[2], last[2]+c)
		if c >= 0 {
			curr[1] = max(curr[1], last[1]+c)
			curr[2] = max(curr[2], last[2]+c)
		} else {
			curr[1] = max(curr[1], last[0])
			curr[2] = max(curr[2], last[1])
		}
		return curr
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := coins[i][j]
			if i != 0 {
				dp[i][j] = helper(dp[i][j], dp[i-1][j], c)
			}
			if j != 0 {
				dp[i][j] = helper(dp[i][j], dp[i][j-1], c)
			}
		}
	}
	return max(dp[m-1][n-1][1], dp[m-1][n-1][2])
}
