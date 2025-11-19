package main

import "sort"

func sortMatrix(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	vals := make([]int, 0, m)
	for row := 0; row < m; row++ {
		for i, j := row, 0; i < m; i, j = i+1, j+1 {
			vals = append(vals, grid[i][j])
		}
		sort.Ints(vals)
		for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
			vals[i], vals[j] = vals[j], vals[i]
		}
		for idx, i, j := 0, row, 0; i < m; i, j, idx = i+1, j+1, idx+1 {
			grid[i][j] = vals[idx]
		}
		vals = vals[:0]
	}
	for col := 1; col < n; col++ {
		for i, j := 0, col; j < n; i, j = i+1, j+1 {
			vals = append(vals, grid[i][j])
		}
		sort.Ints(vals)
		for idx, i, j := 0, 0, col; i < n; i, j, idx = i+1, j+1, idx+1 {
			grid[i][j] = vals[idx]
		}
		vals = vals[:0]
	}
	return grid
}
