package main

import "math"

func mincostTickets(days []int, costs []int) int {
	for i := 1; i >= 0; i-- {
		costs[i] = min(costs[i], costs[i+1])
	}
	dp := make([]int, len(days)+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i, d := range days {
		// 1 day
		dp[i+1] = min(dp[i+1], dp[i]+costs[0])
		// 7 day
		for j := i + 1; j < len(days) && days[j] < d+7; j++ {
			dp[j+1] = min(dp[j+1], dp[i]+costs[1])
		}
		// 30 day
		for j := i + 1; j < len(days) && days[j] < d+30; j++ {
			dp[j+1] = min(dp[j+1], dp[i]+costs[2])
		}
	}
	return dp[len(days)]
}
