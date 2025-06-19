package main

import (
	"sort"
)

func maximizeWin(prizePositions []int, k int) int {
	var ret int
	dp := make([]int, len(prizePositions)+1)
	for i := len(prizePositions) - 1; i >= 0; i-- {
		v := prizePositions[i]
		idx := sort.SearchInts(prizePositions, v+k+1)
		dp[i] = max(idx-i, dp[i+1])
		ret = max(ret, idx-i+dp[idx])
	}
	return ret
}
