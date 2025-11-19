package main

func climbStairs(n int, costs []int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i >= 2 {
			dp[i] = min(dp[i], dp[i-2]+4)
		}
		if i >= 3 {
			dp[i] = min(dp[i], dp[i-3]+9)
		}
		dp[i] += costs[i-1]
	}
	return dp[n]
}
