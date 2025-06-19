package main

func coinChange(coins []int, n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, c := range coins {
			if i < c {
				continue
			}
			if dp[i-c] < 0 {
				continue
			}
			if dp[i] == -1 {
				dp[i] = dp[i-c] + 1
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	return dp[n]
}
