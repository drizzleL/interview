package main

import (
	"math"
)

func minCut(s string) int {
	dp := make([]int, len(s)+1)
	apli := map[[2]int]bool{}
	isPali := func(i, j int) bool {
		if s[i] != s[j] {
			return false
		}
		ret := j-i <= 2
		ret = ret || apli[[2]int{i + 1, j - 1}]
		apli[[2]int{i, j}] = ret
		return ret
	}
	for i := 0; i < len(s); i++ {
		dp[i+1] = math.MaxInt32
		for j := i; j >= 0; j-- {
			if !isPali(j, i) {
				continue
			}
			dp[i+1] = min(dp[i+1], dp[j]+1)
		}
	}
	return dp[len(s)] - 1
}
