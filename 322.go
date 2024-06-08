package main

func coinChange(coins []int, n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for _, c := range coins {
			if i < c {
				continue
			}
			if dp[i-c] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	return dp[n]
}
