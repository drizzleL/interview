package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > i {
				break
			}
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + dp[i][j+1]
			} else {
				dp[i+1][j+1] = dp[i][j+1]
			}
		}
	}
	return dp[m][n]
}
func numDistinct2(s string, t string) int {

	var helper func(i, j int) int
	cache := map[[2]int]int{}
	helper = func(i, j int) (ret int) {
		if c, ok := cache[[2]int{i, j}]; ok {
			return c
		}
		defer func() {
			cache[[2]int{i, j}] = ret
		}()
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		if len(s)-i < len(t)-j {
			return 0
		}
		if s[i] == t[j] {
			ret += helper(i+1, j+1)
		}
		ret += helper(i+1, j)
		return ret
	}
	return helper(0, 0)
}
