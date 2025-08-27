package main

import "math"

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}
	q := [][2]int{{0, -1}}
	var dfs func(i, j int, k int, next *[][2]int)
	dfs = func(i, j int, k int, next *[][2]int) {
		for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			x, y := i+dir[0], j+dir[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if dp[x][y] <= k { // used before
				continue
			}
			if grid[x][y] == 0 {
				dp[x][y] = k
				dfs(x, y, k, next)
				continue
			}
			dp[x][y] = k + 1
			*next = append(*next, [2]int{x, y})
		}
	}
	for k := 0; k < health && len(q) > 0; k++ {
		var next [][2]int
		for _, v := range q {
			dfs(v[0], v[1], k, &next)
		}
		q = next
	}
	return dp[m-1][n-1] < health
}
