package main

func numWays(steps int, arrLen int) int {
	if steps+1 < arrLen {
		arrLen = steps + 1
	}
	dp := make([]int, arrLen)
	dp[0] = 1
	for i := 0; i < steps; i++ {
		next := make([]int, arrLen)
		for i := range next {
			next[i] = dp[i]
			if i > 0 {
				next[i] += dp[i-1]
			}
			if i+1 < arrLen {
				next[i] += dp[i+1]
			}
		}
		dp = next
	}
	return dp[0]
}
