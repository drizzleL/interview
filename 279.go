package main

func numSquares(n int) int {
	candidates := []int{}
	for i := 1; i*i <= n; i++ {
		candidates = append(candidates, i*i)
	}
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for _, c := range candidates {
			if i < c {
				continue
			}
			if dp[i-c] == -1 {
				continue
			}
			if dp[i] == -1 || dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	return dp[n]
}
