package main

func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1
	}
	dp := make([]float64, n+1)
	dp[0] = 1
	s := 1.0
	for i := 1; i <= n; i++ {
		dp[i] = s / float64(maxPts)
		if i < k {
			s += dp[i]
		}
		if i-maxPts >= 0 && i-maxPts < k {
			s -= dp[i-maxPts]
		}
	}
	var ret float64
	for i := k; i <= n; i++ {
		ret += dp[i]
	}
	return ret
}
