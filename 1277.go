package main

func countSquares(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	var ret int
	m, n := len(matrix), len(matrix[0])
	row, col := make([][]int, m), make([][]int, m)
	for i := range row {
		row[i] = make([]int, n)
		col[i] = make([]int, n)
	}
	if matrix[0][0] == 1 {
		row[0][0] = 1
		col[0][0] = 1
		ret += 1
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			continue
		}
		ret += 1
		col[i][0] = col[i-1][0] + 1
		row[i][0] = 1
	}
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			continue
		}
		ret += 1
		row[0][j] = row[0][j-1] + 1
		col[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			row[i][j] = row[i][j-1] + 1
			col[i][j] = col[i-1][j] + 1
			size := matrix[i-1][j-1] + 1
			size = min(row[i][j], size)
			size = min(col[i][j], size)
			ret += size
			matrix[i][j] = size
		}
	}
	return ret
}
