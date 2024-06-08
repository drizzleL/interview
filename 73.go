package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	var row, col bool
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				continue
			}
			if i == 0 {
				row = true
			}
			if j == 0 {
				col = true
			}
			matrix[i][0] = 0
			matrix[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	for i := 0; i < m && col; i++ {
		matrix[i][0] = 0
	}
	for j := 1; j < n && row; j++ {
		matrix[0][j] = 0
	}
}
