package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1); i++ {
		for j := 0; j < len(word2); j++ {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
				continue
			}
			dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
		}
	}
	val := dp[len(word1)][len(word2)]
	return len(word1) - val + len(word2) - val
}
