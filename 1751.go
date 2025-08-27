package main

import "sort"

func maxValue3(events [][]int, k int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, len(events)+1)
	}
	for i := len(events) - 1; i >= 0; i-- {
		dp[0][i] = events[i][2]
	}
	for i := len(events) - 1; i >= 0; i-- {
		nextIdx := sort.Search(len(events), func(x int) bool {
			return events[x][0] > events[i][1]
		})
		for j := 1; j < k; j++ {
			dp[j][i] = dp[j-1][i]
			dp[j][i] = max(dp[j][i], dp[j][i+1])
			dp[j][i] = max(dp[j][i], dp[j-1][nextIdx]+events[i][2])
		}
	}
	return dp[k-1][0]
}
