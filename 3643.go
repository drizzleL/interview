package main

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for row1, row2 := x, x+k-1; row1 < row2; row1, row2 = row1+1, row2-1 {
		for col := y; col < y+k; col++ {
			grid[row1][col], grid[row2][col] = grid[row2][col], grid[row1][col]
		}
	}
	return grid
}
