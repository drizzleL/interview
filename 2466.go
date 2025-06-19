package main

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[zero] += 1
	dp[one] += 1
	var ret int
	for i := 1; i < low; i++ {
		dp[i] += dp[max(0, i-one)]
		dp[i] += dp[max(0, i-zero)]
		dp[i] %= 1e9 + 7
		if i >= low && i <= high {
			ret += dp[i]
			ret %= 1e9 + 7
		}
	}
	return ret
}
