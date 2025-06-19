package main

func waysToReachTarget(target int, types [][]int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for _, t := range types {
		cnt, point := t[0], t[1]
		for i := target; i >= 0; i-- {
			for j := 1; j <= cnt && i-j*point >= 0; j++ {
				dp[i] += dp[i-j*point]
				dp[i] %= 1e9 + 7
			}
		}
	}
	return dp[target]
}
