package main

func numTilings(n int) int {
	dp := make([][3]int, n+1)
	dp[1][0] = 1
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-2][1] + dp[i-2][2] + dp[i-2][0]
		dp[i][0] %= 1e9 + 7
		dp[i][1] = dp[i-1][0]
		dp[i][2] = dp[i-1][0]
	}
	return dp[n][0]
}
