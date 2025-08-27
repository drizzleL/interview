package main

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}
	var sum int
	for i := 0; i < m; i++ {
		sum += grid[i][0]
		sum %= k
		dp[i][0][sum] = 1
	}
	sum = 0
	for j := 0; j < n; j++ {
		sum += grid[0][j]
		sum %= k
		dp[0][j][sum] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for s := 0; s < k; s++ {
				news := (s + grid[i][j]) % k
				dp[i][j][news] += dp[i-1][j][s] + dp[i][j-1][s]
				dp[i][j][news] %= 1e9 + 7
			}
		}
	}
	return dp[m-1][n-1][0]
}
