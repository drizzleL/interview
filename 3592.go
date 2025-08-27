package main

func findCoins(numWays []int) []int {
	var ret []int
	dp := make([]int, len(numWays)+1)
	dp[0] = 1
	copy(dp[1:], numWays)
	for i, v := range dp {
		if v > 1 {
			return nil
		}
		if v == 0 {
			continue
		}
		ret = append(ret, i)
		for j := len(dp) - 1; j >= i; j-- {
			dp[j] -= dp[j-i]
			if dp[j] < 0 {
				return nil
			}
		}
	}
	return ret
}
