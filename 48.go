package main

func rotate(matrix [][]int) {
	size := len(matrix)
	if size == 0 {
		return
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < size; i++ {
		for l, r := 0, size-1; l < r; l, r = l+1, r-1 {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
		}
	}
}
