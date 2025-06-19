package main

func equalPairs(grid [][]int) int {
	var ret int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			k := 0
			for grid[i][k] == grid[k][j] {
				k += 1
			}
			if k == len(grid) {
				ret += 1
			}
		}
	}
	return ret
}
