package main

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		switch {
		case matrix[i][j] == target:
			return true
		case matrix[i][j] > target:
			j -= 1
		case matrix[i][j] < target:
			i += 1
		}
	}
	return false
}
