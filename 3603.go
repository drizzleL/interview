package main

import "math"

func minCost10(m int, n int, waitCost [][]int) int64 {
	cost := make([][]int, m)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			baseCost := (i + 1) * (j + 1)
			if i == 0 && j == 0 {
				cost[i][j] = baseCost + waitCost[i][j]
				continue
			}
			if i != 0 {
				cost[i][j] = min(cost[i][j], cost[i-1][j]+baseCost)
			}
			if j != 0 {
				cost[i][j] = min(cost[i][j], cost[i][j-1]+baseCost)
			}
			if i != m-1 || j != n-1 {
				cost[i][j] += waitCost[i][j]
			}
		}
	}
	return int64(cost[m-1][n-1])
}
