package main

func numberOfRightTriangles(grid [][]int) int64 {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	rows, cols := make([]int, m), make([]int, n)
	var ret int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			rows[i] += 1
			cols[j] += 1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			ret += (rows[i] - 1) * (cols[j] - 1)
		}
	}
	return int64(ret)
}
