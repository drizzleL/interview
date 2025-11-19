package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
	dp := make([]int, forget)
	dp[0] = 1
	var share int
	for i := 1; i < n; i++ {
		if i >= delay {
			share += dp[(i-delay)%forget]
			share %= 1e9 + 7
		}
		if i >= forget {
			share -= dp[i%forget]
			share += 1e9 + 7
			share %= 1e9 + 7
		}
		dp[i%forget] = share
		share %= 1e9 + 7
	}
	var ret int
	for _, v := range dp {
		ret += v
		ret %= 1e9 + 7
	}
	return ret
}
