package main

func waysToChange(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	coins := []int{1, 5, 10, 25}
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[n]
}
