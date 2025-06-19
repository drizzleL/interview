package main

import "sort"

func largestSubmatrix(matrix [][]int) int {
	dp := make([]int, len(matrix[0]))
	var ret int
	helper := func() int {
		cpy := make([]int, len(dp))
		copy(cpy, dp)
		sort.Ints(cpy)
		var ret int
		for i := 0; i < len(cpy); i++ {
			ret = max(ret, cpy[i]*(len(cpy)-i))
		}
		return ret
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				dp[j] = 0
				continue
			}
			dp[j] += 1
		}
		ret = max(ret, helper())
	}
	return ret
}
