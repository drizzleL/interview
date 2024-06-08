package main

func maximizeTheProfit(n int, offers [][]int) int {
	dict := map[int][][]int{}
	for _, offer := range offers {
		dict[offer[1]] = append(dict[offer[1]], offer)
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		for _, offer := range dict[i] {
			dp[i] = max(dp[i], dp[offer[0]]+offer[2])
		}
	}
	return dp[n]
}
