package main

import "log"

func numPermsDISequence(s string) int {
	n := len(s)
	if n == 0 {
		return 1
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 1
	}
	for i, op := range s {
		switch op {
		case 'I':
			for j, curr := 0, 0; j < n-i; j++ {
				curr += dp[i][j]
				dp[i+1][j] = curr
				dp[i+1][j] %= 1e9 + 7
			}
		case 'D':
			for j, curr := n-i-1, 0; j >= 0; j-- {
				curr += dp[i][j+1]
				dp[i+1][j] = curr
				dp[i+1][j] %= 1e9 + 7
			}
		}
	}
	log.Println(dp)
	return dp[n][0]
}
