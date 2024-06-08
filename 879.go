package main

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, minProfit+1)
	}
	dp[0][0] = 1
	for i := range group {
		g1, p1 := group[i], profit[i]
		for p := minProfit; p >= 0; p-- {
			for g := n; g >= g1; g-- {
				dp[g][p] += dp[g-g1][max(0, p-p1)]
				dp[g][p] %= 1e9 + 7
			}
		}
	}
	var ret int
	for p := minProfit; p <= minProfit; p++ {
		for g := 0; g <= n; g++ {
			ret += dp[g][p]
			ret %= 1e9 + 7
		}
	}
	return ret
}
