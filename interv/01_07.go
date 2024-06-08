package main

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	for row := 0; row < m; row++ {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			matrix[row][l], matrix[row][r] = matrix[row][r], matrix[row][l]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n-1 {
				break
			}
			j2 := n - i - 1
			i2 := j2 + j - i

			matrix[i][j], matrix[i2][j2] = matrix[i2][j2], matrix[i][j]
		}
	}
}
