package main

import "math"

func mergeStones(stones []int, k int) int {
	size := len(stones)
	for size >= k {
		size = size/k + size%k
	}
	if size != 1 {
		return -1
	}
	dp := make([][]int, len(stones))
	for i := range dp {
		dp[i] = make([]int, len(stones))
	}
	presums := make([]int, len(stones)+1)
	for i := 1; i < len(presums); i++ {
		presums[i] = presums[i-1] + stones[i-1]
	}
	for size := k; size <= len(stones); size += 1 {
		for i := 0; i+size-1 < len(stones); i++ {
			j := i + size - 1
			dp[i][j] = math.MaxInt32
			for mid := i; mid < j; mid += k - 1 {
				dp[i][j] = min(dp[i][j], dp[i][mid]+dp[mid+1][j])
			}
			if (j-i)%(k-1) == 0 {
				dp[i][j] += presums[j+1] - presums[i]
			}
		}
	}
	return dp[0][len(stones)-1]
}
