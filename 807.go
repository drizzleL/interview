package main

func maxIncreaseKeepingSkyline(grid [][]int) int {
	left := make([]int, len(grid))
	up := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			left[i] = max(left[i], grid[i][j])
			up[j] = max(up[j], grid[i][j])
		}
	}
	var ret int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			gap := min(left[i], up[j])
			ret += gap - grid[i][j]
		}
	}
	return ret
}
