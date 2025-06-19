package main

func longestPalindrome5(s string, t string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(t))
	}
	a, b := make([]int, len(s)), make([]int, len(t))
	helper := func(s string) [][]bool {
		dp := make([][]bool, len(s))
		for i := range dp {
			dp[i] = make([]bool, len(s))
			dp[i][i] = true
		}
		for size := 2; size <= len(s); size++ {
			for i := 0; i+size-1 < len(s); i++ {
				j := i + size - 1
				if s[i] != s[j] {
					continue
				}
				if i+1 == j {
					dp[i][j] = true
					continue
				}
				if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
		}
		return dp
	}
	var ret int
	dp1, dp2 := helper(s), helper(t)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if dp1[i][j] {
				a[i] = max(a[i], j-i+1)
			}
			ret = max(ret, a[i])
		}
	}
	for i := len(t) - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if dp2[j][i] {
				b[i] = max(b[i], i-j+1)
			}
			ret = max(ret, b[i])
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				var extra int
				if i+1 < len(s) {
					extra = max(extra, a[i+1])
				}
				if j-1 >= 0 {
					extra = max(extra, b[j-1])
				}
				if i+1 < len(s) && j-1 >= 0 {
					extra = max(extra, dp[i+1][j-1])
				}
				dp[i][j] = max(dp[i][j], extra+2)
			}
			ret = max(ret, dp[i][j])
		}
	}
	return ret
}
