package main

func lenOfVDiagonal(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][][]int, m)
	for i := range dp {
		dp[i] = make([][][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([][]int, 4)
			for k := range dp[i][j] {
				dp[i][j][k] = make([]int, 2)
				dp[i][j][k][0] = -1
				dp[i][j][k][1] = -1
			}
		}
	}
	dirs := [][2]int{{0, 1}, {1, 1}, {1, 0}, {1, -1}}
	var dfs func(i, j int, dir int, target int, turn int) int
	dfs = func(i, j int, dir int, target int, turn int) int {
		nextI, nextJ := i+dirs[dir][0], j+dirs[dir][1]
		if nextI < 0 || nextI >= m || nextJ < 0 || nextJ >= n {
			return 0
		}
		if grid[nextI][nextJ] != target {
			return 0
		}
		if dp[nextI][nextJ][dir][turn] != -1 {
			return dp[nextI][nextJ][dir][turn]
		}
		ret := dfs(nextI, nextJ, dir, target, turn)
		if turn == 0 {
			ret = max(ret, dfs(nextI, nextJ, (dir+1)%4, target, 1)+1)
		}
		dp[nextI][nextJ][dir][turn] = ret
		return ret
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				continue
			}
			for dir := 0; dir < 4; dir++ {
				dfs(i, j, dir, 2, 0)
			}
		}
	}
	return 0
}
