package main

import "math"

func maxMatrixSum(matrix [][]int) int64 {
	m := len(matrix)
	var flag int
	var sum int
	minVal := math.MaxInt32
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			sum += abs(matrix[i][j])
			if matrix[i][j] < 0 {
				flag += 1
			}
			minVal = min(minVal, abs(matrix[i][j]))
		}
	}
	if flag%2 == 1 {
		sum -= minVal * 2
	}
	return int64(sum)
}
