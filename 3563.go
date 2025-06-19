package main

import "strings"

func lexicographicallySmallestString(s string) string {
	match := func(a, b byte) bool {
		if a > b {
			a, b = b, a
		}
		return b-a == 1 || b-a == 25
	}
	dp := make([][]string, len(s)+1)
	for i := range dp {
		dp[i] = make([]string, len(s)+1)
	}
	for size := 1; size <= len(s); size++ {
		for i := 0; i+size <= len(s); i++ {
			j := i + size
			dp[i][j] = s[i:j]
			for k := i + 1; k < j; k++ {
				if match(s[i], s[k]) && dp[i+1][k] == "" && strings.Compare(dp[k+1][j], dp[i][j]) < 0 {
					dp[i][j] = dp[k+1][j]
				}
			}
		}
	}
	return dp[0][len(s)]
}
