package main

import "sort"

func makesquare(matchsticks []int) bool {
	var sum int
	for _, m := range matchsticks {
		sum += m
	}
	if sum%4 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	edgeLen := sum / 4
	dp := [4]int{}
	var helper func(i int) bool
	helper = func(i int) (ret bool) {
		if i == len(matchsticks) {
			return true
		}
		for j := 0; j < 4; j++ {
			if j > 0 && dp[j] == dp[j-1] {
				continue
			}
			if dp[j]+matchsticks[i] > edgeLen {
				continue
			}
			dp[j] += matchsticks[i]
			if helper(i + 1) {
				return true
			}
			dp[j] -= matchsticks[i]
		}
		return false
	}
	return helper(0)
}
