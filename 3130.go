package main

func numberOfStableArrays(zero int, one int, limit int) int {
	dp := make([][][]int, zero+1)
	for i := range dp {
		dp[i] = make([][]int, one+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i <= limit && i <= zero; i++ {
		dp[i][0][0] = 1
	}
	for i := 0; i <= limit && i <= one; i++ {
		dp[0][i][1] = 1
	}
	for i := 1; i <= zero; i++ {
		for j := 1; j <= one; j++ {
			dp[i][j][0] = dp[i-1][j][0] + dp[i-1][j][1]
			dp[i][j][1] = dp[i][j-1][0] + dp[i][j-1][1]
			if i > limit {
				dp[i][j][0] -= dp[i-limit-1][j][1]
			}
			if j > limit {
				dp[i][j][1] -= dp[i][j-limit-1][0]
			}
			dp[i][j][0] += 1e9 + 7
			dp[i][j][0] %= 1e9 + 7
			dp[i][j][1] += 1e9 + 7
			dp[i][j][0] %= 1e9 + 7
		}
	}
	return (dp[zero][one][0] + dp[zero][one][1]) % (1e9 + 7)
}
