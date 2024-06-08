package main

func findMaxFish(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if grid[i][j] <= 0 {
			return 0
		}
		ret := grid[i][j]
		grid[i][j] = -1
		ret += dfs(i+1, j)
		ret += dfs(i-1, j)
		ret += dfs(i, j+1)
		ret += dfs(i, j-1)
		return ret
	}
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ret = max(ret, dfs(i, j))
		}
	}
	return ret
}
