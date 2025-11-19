package main

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][2]int, len(grid))
	for i := range dp {
		dp[i] = make([][2]int, len(grid[0]))
		for j := range dp[i] {
			dp[i][j] = [2]int{-1, -1}
		}
	}
	var dfs func(i, j int, dir int) int
	dfs = func(i, j int, dir int) int {
		if i >= m || j >= n {
			return 0
		}
		if i == m-1 && j == n-1 {
			return 1
		}
		if dp[i][j][dir] != -1 {
			return dp[i][j][dir]
		}
		dp[i][j][dir] = 0
		if grid[i][j] == 0 {
			dp[i][j][dir] += dfs(i+1, j, 0)
			dp[i][j][dir] += dfs(i, j+1, 1)
		} else {
			if dir == 0 {
				dp[i][j][dir] = dfs(i+1, j, 1)
			} else {
				dp[i][j][dir] = dfs(i, j+1, 0)
			}
		}
		dp[i][j][dir] %= 1e9 + 7
		return dp[i][j][dir]
	}
	ret := dfs(0, 0, 0) + dfs(0, 0, 1)
	ret %= 1e9 + 7
	return ret
}
