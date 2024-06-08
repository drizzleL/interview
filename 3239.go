package main

func minFlips(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var flipRow, flipCol int
	for row := 0; row < m; row++ {
		for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
			if grid[row][i] != grid[row][j] {
				flipRow += 1
			}
		}
	}
	for col := 0; col < n; col++ {
		for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
			if grid[i][col] != grid[j][col] {
				flipCol += 1
			}
		}
	}
	return min(flipRow, flipCol)
}
