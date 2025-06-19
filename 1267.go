package main

func countServers2(grid [][]int) int {
	rows := make([]bool, len(grid))
	cols := make([]bool, len(grid[0]))
	var ret int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}
			if rows[i] || cols[i] {
				ret += 1
			}
			rows[i] = true
			cols[j] = true
		}
	}
	return ret
}
