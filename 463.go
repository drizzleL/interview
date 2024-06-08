package main

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var ret int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			ret += 4
			if i != 0 && grid[i-1][j] == 1 { // up exists
				ret -= 2
			}
			if j != 0 && grid[i][j-1] == 1 {
				ret -= 2
			}
		}
	}
	return ret
}
