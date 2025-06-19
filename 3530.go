package main

import (
	"math/bits"
)

func maxProfit6(n int, edges [][]int, score []int) int {
	need := make([]int, (1<<n)-1)
	for _, ed := range edges {
		need[ed[1]] |= 1 << ed[0]
	}
	dp := make([]int, 1<<n)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	var dfs func(mask int) int // reach mask, max profit
	dfs = func(mask int) (ret int) {
		if dp[mask] != -1 {
			return dp[mask]
		}
		defer func() {
			dp[mask] = ret
		}()
		ret = -2
		pos := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			mask2 := mask ^ 1<<i
			if need[i]|mask2 != mask2 { // need more
				continue
			}
			preRet := dfs(mask2)
			if preRet == -2 {
				continue
			}
			ret = max(ret, preRet+score[i]*pos)
		}
		return
	}
	return dfs((1 << n) - 1)
}
