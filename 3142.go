package main

func satisfiesConditions(grid [][]int) bool {
	for j := 0; j < len(grid[0]); j++ {
		if j != 0 && grid[0][j] == grid[0][j-1] {
			return false
		}
		for i := 1; i < len(grid); i++ {
			if grid[i][j] != grid[0][j] {
				return false
			}
		}
	}
	return true
}
