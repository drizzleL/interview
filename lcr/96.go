package main

func isInterleave(s1 string, s2 string, s3 string) (ret bool) {
	if len(s1) == 0 || len(s2) == 0 {
		return s1+s2 == s3
	}
	if s1[0] != s3[0] && s2[0] != s3[0] {
		return false
	}
	if s2[0] != s3[0] {
		return isInterleave(s1[1:], s2, s3[1:])
	}
	if s1[0] != s3[0] {
		return isInterleave(s1, s2[1:], s3[1:])
	}
	return isInterleave(s1, s2[1:], s3[1:]) || isInterleave(s1[1:], s2, s3[1:])
}

func isInterleave2(s1 string, s2 string, s3 string) (ret bool) {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return s1+s2 == s3
	}
	if m+n != len(s3) {
		return false
	}
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j] && s3[i+j-1] == s1[i-1]
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || dp[i][j-1] && s3[i+j-1] == s2[j-1]
			}
		}
	}
	return dp[m][n]
}
