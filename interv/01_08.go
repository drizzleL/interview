package main

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	var row0, col0 bool
	for col := 0; col < n; col++ {
		if matrix[0][col] == 0 {
			row0 = true
			break
		}
	}
	for row := 0; row < m; row++ {
		if matrix[row][0] == 0 {
			col0 = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if matrix[0][col] != 0 && matrix[row][0] != 0 {
				continue
			}
			matrix[row][col] = 0
		}
	}
	if row0 {
		for col := 0; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if col0 {
		for row := 0; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}
