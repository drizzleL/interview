package main

func hitBricks(grid [][]int, hits [][]int) []int {
	m, n := len(grid), len(grid[0])
	g2 := make([][]int, m)
	for i := range g2 {
		g2[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if g2[i][j] == 1 { // visited before
			return 0
		}
		if grid[i][j] != 1 {
			return 0
		}
		ret := 1
		g2[i][j] = 1
		ret += dfs(i-1, j)
		ret += dfs(i+1, j)
		ret += dfs(i, j-1)
		ret += dfs(i, j+1)
		return ret
	}
	for _, h := range hits {
		x, y := h[0], h[1]
		if grid[x][y] == 0 {
			continue
		}
		grid[h[0]][h[1]] = 2
	}
	var sum int
	for j := 0; j < n; j++ {
		sum += dfs(0, j)
	}
	ret := make([]int, len(hits))
	checkStable := func(x, y int) bool {
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		return g2[x][y] == 1
	}
	for i := len(hits) - 1; i >= 0; i-- {
		x, y := hits[i][0], hits[i][1]
		if grid[x][y] == 0 {
			continue
		}
		grid[x][y] = 1
		if !(x == 0 || checkStable(x-1, y) || checkStable(x, y-1) || checkStable(x+1, y) || checkStable(x, y+1)) {
			continue
		}
		oldSum := sum
		sum += dfs(x, y)
		ret[i] = sum - oldSum - 1
	}
	return ret
}
