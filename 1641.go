package main

func countVowelStrings(n int) int {
	dp := make([]int, 5)
	for i := range dp {
		dp[i] = 1
	}
	sum := 5
	for j := 1; j < n; j++ {
		var newsum int
		for i := 0; i < len(dp); i++ {
			dp[i], sum = dp[i]+sum, sum-dp[i]
			newsum += dp[i]
		}
		sum = newsum
	}

	return sum
}
