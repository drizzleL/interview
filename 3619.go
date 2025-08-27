package main

func countIslands(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		if grid[i][j] <= 0 {
			return 0
		}
		grid[i][j] = -1
		return grid[i][j] + dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1)
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] <= 0 {
				continue
			}
			v := dfs(i, j)
			if v%k == 0 {
				ret += 1
			}
		}
	}
	return ret
}
