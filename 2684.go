package main

func maxMoves(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	steps := make([][]int, len(grid))
	for i := range steps {
		steps[i] = make([]int, len(grid[0]))
	}
	var ret int
	for j := 1; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			steps[i][j] = -1
			if i != 0 && grid[i-1][j-1] < grid[i][j] && steps[i-1][j-1] != -1 {
				steps[i][j] = max(steps[i][j], steps[i-1][j-1]+1)
			}
			if grid[i][j-1] < grid[i][j] && steps[i][j-1] != -1 {
				steps[i][j] = max(steps[i][j], steps[i][j-1]+1)
			}
			if i != len(grid)-1 && grid[i+1][j-1] < grid[i][j] && steps[i+1][j-1] != -1 {
				steps[i][j] = max(steps[i][j], steps[i+1][j-1]+1)
			}
			ret = max(ret, steps[i][j])
		}
	}
	return ret
}
