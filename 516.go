package main

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for j := 0; j < len(s); j++ {
		dp[j][j] = 1
		for i := j - 1; i >= 0; i-- {
			dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			}
		}
	}
	return dp[0][len(s)-1]
}
