package main

func findChampion2(grid [][]int) int {
	var ret, wins int
	for i := 0; i < len(grid); i++ {
		var w int
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			w += 1
		}
		if w > wins {
			wins = w
			ret = i
		}
	}
	return ret
}
