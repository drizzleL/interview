package main

func longestCommonSubsequence(text1 string, text2 string) int {
	var helper func(i, j int) int
	cache := map[[2]int]int{}
	helper = func(i, j int) (c int) {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if c, ok := cache[[2]int{i, j}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, j}] = c
		}()
		if text1[i] == text2[j] {
			return helper(i+1, j+1) + 1
		}
		return max(helper(i+1, j), helper(i, j+1))
	}
	return helper(0, 0)
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}
