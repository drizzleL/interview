package main

func minCost3(grid [][]int) int {
	dirs := [][2]int{
		{0, 0},
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	m, n := len(grid), len(grid[0])
	cache := make([]int, m*n)
	for i := range cache {
		cache[i] = -1
	}
	var nodes [][2]int
	var dfs func(x, y int, val int)
	dfs = func(x, y int, val int) {
		if x < 0 || y < 0 || x >= m || y >= n {
			return
		}
		key := x*n + y
		if cache[key] != -1 {
			return
		}
		cache[key] = val
		nodes = append(nodes, [2]int{x, y})
		dir := dirs[grid[x][y]]
		x2, y2 := x+dir[0], y+dir[1]
		dfs(x2, y2, val)
	}
	dfs(0, 0, 0)
	for k := 1; cache[m*n-1] == -1; k++ {
		var prev [][2]int
		nodes, prev = nil, nodes
		for _, node := range prev {
			x, y := node[0], node[1]
			for i := 1; i <= 4; i++ {
				x2, y2 := x+dirs[i][0], y+dirs[i][1]
				dfs(x2, y2, k)
			}
		}
	}
	return cache[m*n-1]
}
