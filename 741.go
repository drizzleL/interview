package main

import (
	"math"
)

func cherryPickup(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == -1 || grid[n-1][n-1] == -1 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	dp[0][0] = grid[0][0]
	for step := 1; step <= n*2-2; step++ {
		newDp := make([][]int, n)
		for i := range newDp {
			newDp[i] = make([]int, n)
			for j := range newDp[i] {
				newDp[i][j] = math.MinInt32
			}
		}
		for r1 := 0; r1 < n; r1++ {
			for r2 := 0; r2 < n; r2++ {
				c1 := step - r1
				c2 := step - r2
				if c1 < 0 || c2 < 0 || c1 >= n || c2 >= n {
					continue
				}
				if grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}
				c := grid[r1][c1]
				if r1 != r2 {
					c += grid[r2][c2]
				}
				for _, d := range [][2]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}} {
					pr1, pr2 := r1-d[0], r2-d[1]
					pc1, pc2 := step-1-pr1, step-1-pr2
					if pr1 < 0 || pr2 < 0 || pc1 >= n || pc2 >= n {
						continue
					}
					if dp[pr1][pr2] == math.MinInt32 {
						continue
					}
					newDp[r1][r2] = max(newDp[r1][r2], dp[pr1][pr2]+c)
				}
			}
		}
		dp = newDp
	}
	if dp[n-1][n-1] == math.MinInt32 {
		return 0
	}
	return dp[n-1][n-1]
}
