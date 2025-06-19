package main

func minimumOperations3(grid [][]int) int {
	var ret int
	for j := 0; j < len(grid[0]); j++ {
		for i := 1; i < len(grid); i++ {
			diff := max(0, grid[i-1][j]+1-grid[i][j])
			ret += diff
			grid[i][j] += diff
		}
	}
	return ret
}
