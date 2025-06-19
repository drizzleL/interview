package main

import "sort"

func maxTaxiEarnings(n int, rides [][]int) int64 {
	sort.Slice(rides, func(i, j int) bool {
		return rides[i][0] < rides[j][0]
	})
	dp := make([]int, len(rides)+1)
	for i := len(rides) - 1; i >= 0; i-- {
		idx := sort.Search(len(rides), func(j int) bool {
			return rides[j][0] >= rides[i][1]
		})
		dp[i] = dp[i+1]
		dp[i] = max(dp[i], dp[idx]+(rides[i][1]-rides[i][0]+rides[i][2]))
	}
	return int64(dp[0])
}
