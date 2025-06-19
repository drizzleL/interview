package main

func mostPoints(questions [][]int) int64 {
	dp := make([]int, len(questions)+1)
	for i := len(questions) - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		var extra int
		if questions[i][1]+i+1 < len(dp) {
			extra = dp[questions[i][1]+i+1]
		}
		dp[i] = max(dp[i], questions[i][0]+extra)
	}
	return int64(dp[0])
}
