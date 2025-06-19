package main

func longestPalindromicSubsequence(s string, k int) int {
	dp := make([][][]int, len(s))
	for i := range dp {
		dp[i] = make([][]int, len(s))
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}
	diff := func(i, j int) int {
		a, b := int(s[i]-'a'), int(s[j]-'a')
		if a > b {
			a, b = b, a
		}
		return min(a+26-b, b-a)
	}
	for j := 0; j < len(s); j++ {
		for m := 0; m <= k; m++ {
			dp[j][j][m] = 1
		}
		for i := j - 1; i >= 0; i-- {
			for m := 0; m <= k; m++ {
				dp[i][j][m] = max(dp[i+1][j][m], dp[i][j-1][m])
				gap := diff(j, i)
				if gap <= m {
					dp[i][j][m] = max(dp[i][j][m], dp[i+1][j-1][m-gap]+2)
				}
			}
		}
	}
	return dp[0][len(s)-1][k]
}
