package main

func matrixScore(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			continue
		}
		for j := 0; j < n; j++ {
			grid[i][j] = 1 - grid[i][j]
		}
	}
	var ret int
	for k := 0; k < n; k++ {
		j := n - k - 1
		var zeros int
		for i := 0; i < m; i++ {
			if grid[i][j] == 0 {
				zeros += 1
			}
		}
		ret += (1 << k) * max(zeros, m-zeros)
	}
	return ret
}
