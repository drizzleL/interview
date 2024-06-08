package main

import (
	"math"
	"regexp"
)

func strangePrinter(s string) int {
	re := regexp.MustCompile(`(.)\\1+`)
	s = re.ReplaceAllString(s, "$1")
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}
	for size := 2; size <= len(s); size++ {
		for i := 0; i <= len(s)-size; i++ {
			j := i + size - 1
			for k := i; k < j; k++ {
				tmp := dp[i][k] + dp[k+1][j]
				if s[i] == s[j] {
					tmp -= 1
				}
				dp[i][j] = min(dp[i][j], tmp)
			}
		}
	}
	return dp[0][len(s)-1]
}
