package main

func maximumProfit(prices []int, k int) int64 {
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, len(prices))
	}
	for i := 1; i <= k; i++ {
		minTmp, maxTmp := prices[0], prices[0]
		for j := 1; j < len(prices); j++ {
			dp[i][j] = max(dp[i][j-1], max(prices[j]-minTmp, maxTmp-prices[j]))
			minTmp = min(minTmp, prices[j]-dp[i-1][j-1])
			maxTmp = max(maxTmp, dp[i-1][j-1]+prices[j])
		}
	}
	return int64(dp[k][len(prices)-1])
}
