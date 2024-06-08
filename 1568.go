package main

func minDays2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int, val int) int
	dfs = func(i, j int, val int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if grid[i][j] == 0 || grid[i][j] == val {
			return 0
		}
		grid[i][j] = val
		ret := 1
		ret += dfs(i-1, j, val)
		ret += dfs(i+1, j, val)
		ret += dfs(i, j+1, val)
		ret += dfs(i, j-1, val)
		return ret
	}
	var cnt int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			if cnt != 0 {
				return 0
			}
			cnt = dfs(i, j, 2)
		}
	}
	if cnt == 0 {
		return 0
	}
	if cnt <= 2 {
		return cnt
	}
	check := func(i, j int, val int) bool {
		c := dfs(i, j, val)
		if c == 0 || c == cnt-1 {
			return false
		}
		return true
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			val := grid[i][j] + 1
			grid[i][j] = val
			if check(i+1, j, val) || check(i-1, j, val) || check(i, j+1, val) || check(i, j-1, val) {
				return 1
			}
		}
	}
	return 2
}
