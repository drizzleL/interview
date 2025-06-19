package main

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	dp := make([][]int, numCarpets+1)
	for i := range dp {
		dp[i] = make([]int, len(floor))
	}
	for i := 0; i < len(floor); i++ {
		if i != 0 {
			dp[0][i] = dp[0][i-1]
		}
		dp[0][i] += int(floor[i] - '0')
	}
	for i := 1; i <= numCarpets; i++ {
		for j := carpetLen; j < len(floor); j++ {
			dp[i][j] = dp[i-1][j-carpetLen]
			if j != 0 {
				dp[i][j] = min(dp[i][j], int(floor[j]-'0')+dp[i][j-1])
			}
		}
	}
	return dp[numCarpets][len(floor)-1]
}
