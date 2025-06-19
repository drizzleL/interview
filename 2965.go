package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	ret := make([]int, 2)
	seen := make([]bool, len(grid)*len(grid)+1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			v := grid[i][j]
			if seen[v] {
				ret[0] = v
			}
			seen[v] = true
		}
	}
	for i := 1; i < len(seen); i++ {
		if !seen[i] {
			ret[1] = i
			break
		}
	}
	return ret
}
