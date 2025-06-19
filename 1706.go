package main

func findBall(grid [][]int) []int {
	ret := make([]int, len(grid[0]))
	for j := 0; j < len(grid[0]); j++ {
		ret[j] = j
		k := j
		for i := 0; i < len(grid); i++ {
			if k+grid[i][k] < 0 || k+grid[i][k] >= len(grid[0]) {
				ret[j] = -1
				break
			}
			if grid[i][k] == -1 && k-1 >= 0 && grid[i][k-1] == grid[i][k] {
			} else if grid[i][k] == 1 && k+1 < len(grid[0]) && grid[i][k] == grid[i][k+1] {
			} else {
				ret[j] = -1
				break
			}
			k += grid[i][k]
			ret[j] = k
		}
	}
	return ret
}
